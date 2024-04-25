import Post from "./Post";

interface Bookmark {
  id: number;
  photo_id: number;
  user_id: number;
}

export interface BookmarkWithRelation {
  id: number;
  photo_id: number;
  user_id: number;
  photo: Post;
  // user: User;
}

export default Bookmark;
