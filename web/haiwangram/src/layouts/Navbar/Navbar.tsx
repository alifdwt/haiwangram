import {
  Avatar,
  Divider,
  Flex,
  HStack,
  Link,
  Stack,
  Text,
} from "@chakra-ui/react";
import { HomeIcon, SearchIcon, User2Icon, UserIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

const navbarList = [
  {
    name: "Beranda",
    path: "/",
    icon: <HomeIcon />,
  },
  {
    name: "Cari",
    path: "/search",
    icon: <SearchIcon />,
  },
  {
    name: "Pengikut",
    path: "/follow",
    icon: <User2Icon />,
  },
  {
    name: "Profil",
    path: "/profile",
    icon: <UserIcon />,
  },
];

const followedPages = [
  {
    name: "Waktunya Memelihara Kucing",
    path: "/waktunya-memelihara-kucing",
  },
  {
    name: "Anjingku Pahlawanku",
    path: "/anjingku-pahlawanku",
  },
];

export default function Navbar() {
  const navigate = useNavigate();
  const currentUrl = location.pathname;

  return (
    <Stack
      bg={"white"}
      _dark={{ bg: "gray.700" }}
      p={3}
      borderRadius={"lg"}
      gap={4}
    >
      <Stack>
        {navbarList.map((item) => (
          <Link
            key={item.name}
            onClick={() =>
              item.name === "Profil"
                ? navigate(`/${item.path}`)
                : navigate(item.path)
            }
            _hover={{
              textDecoration: "none",
              bg: "primary.700",
              color: "white",
            }}
            borderRadius={"lg"}
            py={3}
            bg={currentUrl === item.path ? "primary.700" : "transparent"}
            color={currentUrl === item.path ? "white" : "black"}
            _dark={{ color: "white" }}
          >
            <HStack ml={4}>
              <Text>{item.icon}</Text>
              <Text>{item.name}</Text>
            </HStack>
          </Link>
        ))}
      </Stack>
      <Divider />
      <Stack>
        <Text
          fontWeight={"bold"}
          color={"gray.500"}
          textTransform={"uppercase"}
          letterSpacing={"wider"}
        >
          Halaman yang kamu sukai
        </Text>
        {followedPages.map((item) => (
          <Link
            key={item.name}
            href={`/pages/${item.path}`}
            as={Flex}
            alignItems={"center"}
            gap={2}
          >
            <Avatar
              name={item.name}
              // src={`https://randomuser.me/api/portraits/men/${Math.floor(
              //   Math.random() * 10 + 1
              // )}.jpg`}
              size={"sm"}
              bg={"primary.700"}
            />
            <Text fontWeight={"bold"} noOfLines={1}>
              {item.name}
            </Text>
          </Link>
        ))}
      </Stack>
    </Stack>
  );
}
