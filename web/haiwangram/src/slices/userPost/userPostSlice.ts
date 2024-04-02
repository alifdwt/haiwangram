import Post from "@/interface/Post";
import { createSlice } from "@reduxjs/toolkit";

interface UserPostSlice {
  userPosts: Post[];
}

const initialState: UserPostSlice = {
  userPosts: [],
};

export const userPostSlice = createSlice({
  name: "userPost",
  initialState,
  reducers: {
    getUserPosts: (state, action) => {
      state.userPosts = action.payload;
    },
  },
});

export const { getUserPosts } = userPostSlice.actions;

export default userPostSlice.reducer;
