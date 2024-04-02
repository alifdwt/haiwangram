interface User {
  id: number;
  email: string;
  username: string;
  full_name: string;
  birth_date: string;
  profile_image_url: string;
  password?: string;
  description: string;
  // banner_profile?: string;
  // followers?: any;
  // following?: any;
}

export default User;
