// {
//     "id": 1,
//     "message": "lorem ipsum dolor sit amet consectetur adipiscing",
//     "photo_id": 1,
//     "user_id": 3,
//     "user": {
//       "id": 3,
//       "email": "sqnewm-hkykbr@example.com",
//       "username": "sqnewm-hkykbr",
//       "full_name": "Sqnewm Hkykbr",
//       "profile_image_url": "https://randomuser.me/api/portraits/women/87.jpg",
//       "description": "lorem ipsum dolor sit amet consectetur adipiscing"
//     },
//     "photo": {
//       "id": 1,
//       "caption": "lorem ipsum dolor sit amet consectetur adipiscing elit condimentum duis varius nisl nostra et suspendisse",
//       "title": "lorem ipsum dolor sit amet consectetur adipiscing",
//       "photo_url": "https://picsum.photos/id/34/1280/720",
//       "user_id": 3
//     },
//     "replies": [
//       {
//         "id": 1,
//         "message": "lorem ipsum dolor sit amet consectetur adipiscing elit condimentum duis",
//         "comment_id": 1,
//         "user_id": 3
//       }
//     ]
//   }

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
}

export default Comment;
