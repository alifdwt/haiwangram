import Post from "@/interface/Post";
import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import PostsTab from "./tabs/Posts";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/store";
import { useInfiniteQuery } from "@tanstack/react-query";
import axiosFetch from "@/config/axiosFetch";
import User from "@/interface/User";
import { getUserPosts } from "@/slices/userPost/userPostSlice";

type ProfileTabProps = {
  userData: User;
  userPosts: Post[];
};

export default function ProfileTabs({ userData, userPosts }: ProfileTabProps) {
  const { user } = useSelector((state: RootState) => state.user);
  const dispatch = useDispatch();

  // @ts-expect-error next-line
  const { fetchNextPage } = useInfiniteQuery({
    queryKey: ["userPosts", userData.id],
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

  return (
    <Tabs isFitted>
      <TabList mb={"1em"}>
        <Tab>Posts</Tab>
        <Tab>Likes</Tab>
        {user?.id === userData.id && <Tab>Bookmarks</Tab>}
      </TabList>
      <TabPanels>
        <TabPanel>
          <PostsTab posts={userPosts} postsNextPageButton={fetchNextPage} />
        </TabPanel>
        <TabPanel>Like</TabPanel>
        {user?.id === userData.id && <TabPanel>Bookmark</TabPanel>}
      </TabPanels>
    </Tabs>
  );
}
