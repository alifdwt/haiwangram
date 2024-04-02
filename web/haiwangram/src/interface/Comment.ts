import CommentReply from "./CommentReply";
import Post from "./Post";
import User from "./User";

interface Comment {
  id: number;
  message: string;
  photo_id: number;
  user_id: number;
  user?: User;
  photo?: Post;
  replies?: CommentReply[];
  created_at: string;
  updated_at: string;
}

export default Comment;
