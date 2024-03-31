import { createSlice } from '@reduxjs/toolkit'

const initialState = {
  id: '',
  email: '',
  subreddits: [],
}

export const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUserId: (state, action) => {
      state.id = action.payload
    },
    setEmail: (state, action) => {
      state.email = action.payload
    },
    setSubreddits: (state, action) => {
      state.subreddits = action.payload
    },
  },
})

export const { setUserId, setEmail, setSubreddits } = userSlice.actions

export default userSlice.reducer
