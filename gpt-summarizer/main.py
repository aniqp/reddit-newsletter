from openai import OpenAI
import os
from dotenv import load_dotenv
import json
from datetime import date, datetime
import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore
import redis
from helpers import *

if __name__ == "__main__":
    load_dotenv()
    api_key = os.environ.get("OPENAI_API_KEY")
    model = os.environ.get("OPENAI_MODEL")
    client = OpenAI(api_key=api_key)
    cred = credentials.Certificate("../reddit-newsletter-firebase-key.json")
    app = firebase_admin.initialize_app(cred)
    db = firestore.client()
    print("application started")
    r = redis.Redis(host="localhost", port=6379, decode_responses=True)
    consume_queue(r, client, model, db)
