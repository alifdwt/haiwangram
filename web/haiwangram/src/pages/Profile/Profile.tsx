import useFetchWithId from "@/hooks/useFetchWithId";
import { Box, Skeleton } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import ProfileTopBar from "./components/ProfileTopBar";
import User from "@/interface/User";
import ProfileHeader from "./components/ProfileHeader";
import ProfileTabs from "./components/ProfileTabs";
import { useSelector } from "react-redux";
import { RootState } from "@/store";

export default function Profile() {
  const { username } = useParams();
  const { userPosts } = useSelector((state: RootState) => state.userPost);
  const { data: fetchedUserData, isLoading: isLoadingUserData } =
    useFetchWithId({
      queryKey: "user",
      id: username!,
      fetchRoute: "users",
    });

  const userData: User = fetchedUserData;

  if (isLoadingUserData) {
    return (
      <Box>
        <Skeleton height="20px" />
      </Box>
    );
  }

  return (
    <Box>
      <ProfileTopBar
        fullName={userData.full_name}
        postCount={userPosts.length}
      />
      <ProfileHeader user={userData} />
      <ProfileTabs userData={userData} userPosts={userPosts} />
    </Box>
  );
}
