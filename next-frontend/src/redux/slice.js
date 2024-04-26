import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  id: '',
  accessToken: '',
  email: '',
  subreddits: [],
  joined: ''
}

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUserData: (state, action) => {
      state.id = action.payload.id,
      state.accessToken = action.payload.accessToken
    },
    setEmail: (state, action) => {
      state.email = action.payload
    },
    setSubreddits: (state, action) => {
      state.subreddits = action.payload
    },
    setJoinedDate: (state, action) => {
      state.joined = action.payload
    },
    updateStarredSubreddits: (state, action) => {
      console.log(action.payload)
      state.subreddits = state.subreddits.map(subreddit => {
        return { ...subreddit, starred: action.payload }
      })
    },
    updateStarredSubreddit: (state, action) => {
      state.subreddits = state.subreddits.map(subreddit => {
        if (subreddit.id === action.payload.subredditId) {
          return { ...subreddit, starred: action.payload.starred }
        }
        return subreddit
      })
    }
  },
})

export const { setUserData, setEmail, setJoinedDate, setSubreddits } = userSlice.actions

export default userSlice.reducer
