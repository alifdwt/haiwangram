import User from "@/interface/User";
import { createSlice } from "@reduxjs/toolkit";

interface UserSliceState {
  user?: User;
  accessToken?: string;
}

const initialState: UserSliceState = {
  user: {},
  accessToken: "",
};

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    login: (state, action) => {
      state.accessToken = action.payload;
    },
    getUser: (state, action) => {
      state.user = action.payload;
    },
    logout: (state) => {
      state.accessToken = "";
      state.user = {};
    },
  },
});

export const { login, getUser, logout } = userSlice.actions;

export default userSlice.reducer;
