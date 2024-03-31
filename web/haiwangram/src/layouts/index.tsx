import { Box, Flex } from "@chakra-ui/react";
import Navbar from "./Navbar/Navbar";
import Timeline from "./Timeline/Timeline";
import { Outlet } from "react-router-dom";
import ProfileCard from "./ProfileCard/ProfileCard";
import Topbar from "./Topbar/Topbar";
import Messages from "./Sidebar/Sidebar";

export default function Layout() {
  return (
    <Box
      h="100vh"
      maxH={"100vh"}
      bg={"gray.100"}
      _dark={{ bg: "gray.900" }}
      overflow={"hidden"}
    >
      <Topbar />
      <Flex>
        <Box flex={1} p={4}>
          <ProfileCard />
          <Navbar />
        </Box>
        <Box
          flex={2}
          p={4}
          h={"92vh"}
          overflowX={"hidden"}
          overflowY={"scroll"}
          css={{ "&::-webkit-scrollbar": { display: "none" } }}
        >
          <Timeline>
            <Outlet />
          </Timeline>
        </Box>
        <Box flex={1} p={4}>
          <Messages />
        </Box>
      </Flex>
    </Box>
  );
}
