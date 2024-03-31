import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  account: {
    reddit: null,
    email: null,
  },
  subreddits: [],
}

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser: (state, action) => {
      state.account = action.payload
    },
    setSubreddits: (state, action) => {
      state.subreddits = action.payload
    },
  },
})

export const { setUser, setSubreddits } = userSlice.actions

export default userSlice.reducer
