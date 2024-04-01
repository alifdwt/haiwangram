// {
//     "id": 1,
//     "message": "lorem ipsum dolor sit amet consectetur adipiscing elit condimentum duis",
//     "comment_id": 1,
//     "user_id": 3,
//     "comment": {
//       "id": 1,
//       "message": "lorem ipsum dolor sit amet consectetur adipiscing",
//       "photo_id": 1,
//       "user_id": 3
//     },
//     "user": {
//       "id": 3,
//       "email": "sqnewm-hkykbr@example.com",
//       "username": "sqnewm-hkykbr",
//       "full_name": "Sqnewm Hkykbr",
//       "profile_image_url": "https://randomuser.me/api/portraits/women/87.jpg",
//       "description": "lorem ipsum dolor sit amet consectetur adipiscing"
//     }
//   }

import Comment from "./Comment";
import User from "./User";

interface CommentReply {
  id: number;
  message: string;
  comment_id: number;
  user_id: number;
  comment: Comment;
  user: User;
}

export default CommentReply;
