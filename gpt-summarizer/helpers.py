from openai import OpenAI
import os
from dotenv import load_dotenv
import json
from datetime import date, datetime
import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore
import redis


def process_data(data: json) -> str:
    hot_posts = data["hotposts"]
    gpt_input = ""
    for hot_post in hot_posts:
        title = hot_post["title"]
        body = hot_post["selftext"]
        comments_arr = hot_post["comments"]["data"]["children"]
        comments = ""
        for comment in comments_arr:
            comment_text = comment["data"]["body"]
            comments += comment_text + ","
        formatted_hot_post = (
            "Post title: title"
            + "\n"
            + " Post body: "
            + body
            + "\n"
            + "Post comments: "
            + comments
        )
        gpt_input += formatted_hot_post + "\n"
        print(f"gpt input: {gpt_input}")
    return gpt_input


def get_gpt_output(
    client: OpenAI, model: str, gpt_input: str, subreddit: str
) -> dict[str:str]:
    completion = client.chat.completions.create(
        model=model,
        messages=[
            {
                "role": "assistant",
                "content": "You are a daily newsletter writer that summarizes the past day's hot posts on a subreddit, with a chill, natural tone.",
            },
            {
                "role": "user",
                "content": f'Provide me a newsletter title and body for the following subreddit\'s hot posts and comments, without an intro and signoff. Format response by including "Newsletter Title:" and "Newsletter Body:" in your response. You will receive a selection of hot posts and comments for each post, DO NOT provide a summary per post. I want you to give me a summary of all the posts, like an overview newsletter.: {gpt_input}.',
            },
        ],
        temperature=0,
        max_tokens=500,
    )
    raw_text = completion.choices[0].message.content
    title_start = raw_text.index("Newsletter Title:") + len("Newsletter Title:")
    title_end = raw_text.index("Newsletter Body:")
    title = raw_text[title_start:title_end].strip()
    body_start = raw_text.index("Newsletter Body:") + len("Newsletter Body:")
    body = raw_text[body_start:].strip()
    today = date.today()
    firestore_today = datetime(today.year, today.month, today.day)
    return {
        "title": title,
        "body": body,
        "subreddit": subreddit,
        "date": firestore_today,
    }


def upload_summary_to_firebase(db, gpt_dict):
    subreddit = gpt_dict["subreddit"]
    subreddit_doc_name = subreddit[2:]
    doc_path = f"daily_subreddit_summary/{subreddit_doc_name}"
    doc_ref = db.document(doc_path)
    if doc_ref.get().exists == False:
        subreddit_path = f"daily_subreddit_summary/{subreddit_doc_name}"
        subreddit_doc = db.document(subreddit_path)
        subreddit_doc.set({})
    dates_subcollection_path = f"daily_subreddit_summary/{subreddit_doc_name}/dates"
    summary_doc_ref = db.collection(dates_subcollection_path).document(
        f"{date.today()}"
    )
    summary_doc_ref.set(gpt_dict)


def consume_queue(
    r: redis.Redis, client: OpenAI, model: str, cred: credentials.Certificate
):
    while True:
        result = r.brpop(keys="hotPostsQueue", timeout=0)
        json_data = result[1]
        data_dict = json.loads(json_data)
        subreddit_name = data_dict["subredditname"]
        gpt_input = process_data(data_dict)
        gpt_output = get_gpt_output(client, model, gpt_input, subreddit_name)
        # print(f"Subreddit name: {subreddit_name}, GPT Output: {gpt_output}")
        upload_summary_to_firebase(cred, gpt_output)
