import Bookmark from "./Bookmark";
import Comment from "./Comment";
import Like from "./Like";
import User from "./User";

interface Post {
  id: number;
  caption: string;
  title: string;
  photo_url: string;
  user_id: number;
  created_at: string;
  updated_at: string;
  user?: User;
  comments?: Comment[] | null;
  likes?: Like[] | null;
  bookmarks?: Bookmark[] | null;
}

export default Post;
