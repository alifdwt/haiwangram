import { persistStore, persistReducer } from "redux-persist";
import { combineReducers, configureStore } from "@reduxjs/toolkit";
import UserReducer from "@/slices/user/userSlice";
import PostReducer from "@/slices/posts/postSlice";
import storage from "redux-persist/lib/storage";
import UserPostReducer from "@/slices/userPost/userPostSlice";
import BookmarkReducer from "@/slices/bookmark/bookmarkSlice";

const persistConfig = {
  key: "root",
  storage,
};

const rootReducer = combineReducers({
  user: UserReducer,
  post: PostReducer,
  userPost: UserPostReducer,
  bookmark: BookmarkReducer,
});

const persistedReducer = persistReducer(persistConfig, rootReducer);

export const store = configureStore({
  reducer: persistedReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: false,
    }),
});

export const persistor = persistStore(store);
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
