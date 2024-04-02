import useFetchWithId from "@/hooks/useFetchWithId";
import { Box, Skeleton } from "@chakra-ui/react";
import { useParams } from "react-router-dom";
import ProfileTopBar from "./components/ProfileTopBar";
import User from "@/interface/User";
import ProfileHeader from "./components/ProfileHeader";
import ProfileTabs from "./components/ProfileTabs";
import { useInfiniteQuery } from "@tanstack/react-query";
import axiosFetch from "@/config/axiosFetch";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/store";
import { getUserPosts } from "@/slices/userPost/userPostSlice";

export default function Profile() {
  const dispatch = useDispatch();
  const { userPosts } = useSelector((state: RootState) => state.userPost);
  const { username } = useParams();
  const { data: fetchedUserData, isLoading: isLoadingUserData } =
    useFetchWithId({
      queryKey: "user",
      id: username!,
      fetchRoute: "users",
    });

  const userData: User = fetchedUserData;

  // @ts-expect-error next-line
  const { fetchNextPage } = useInfiniteQuery({
    queryKey: ["userPosts", username],
    queryFn: async ({ pageParam = 1 }) => {
      const { data } = await axiosFetch.get(
        `/photos/?limit=${pageParam * 4}&userId=${userData.id}`
      );
      dispatch(getUserPosts(data));
      return data;
    },
    getNextPageParam: (_, pages) => {
      return pages.length + 1;
    },
    initialData: {
      pages: [userPosts],
      pageParams: [1],
    },
  });

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
      <ProfileTabs posts={userPosts} postsNextPageButton={fetchNextPage} />
    </Box>
  );
}
