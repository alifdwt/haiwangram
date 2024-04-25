import { BookmarkWithRelation } from "@/interface/Bookmark";
import { createSlice } from "@reduxjs/toolkit";

interface BookmarkSliceState {
  bookmarks: BookmarkWithRelation[];
}

const initialState: BookmarkSliceState = {
  bookmarks: [],
};

export const bookmarkSlice = createSlice({
  name: "bookmark",
  initialState,
  reducers: {
    getBookmarks: (state, action) => {
      state.bookmarks = action.payload;
    },
  },
});

export const { getBookmarks } = bookmarkSlice.actions;

export default bookmarkSlice.reducer;
