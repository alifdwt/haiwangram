import Post from "@/interface/Post";
import { createSlice } from "@reduxjs/toolkit";

interface PostSliceState {
  posts: Post[];
}

const initialState: PostSliceState = {
  posts: [],
};

export const postSlice = createSlice({
  name: "post",
  initialState,
  reducers: {
    getPosts: (state, action) => {
      state.posts = action.payload;
    },
  },
});

export const { getPosts } = postSlice.actions;

export default postSlice.reducer;
