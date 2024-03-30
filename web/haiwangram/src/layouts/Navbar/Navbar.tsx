import { HStack, Link, Stack, Text } from "@chakra-ui/react";
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

export default function Navbar() {
  const navigate = useNavigate();
  const currentUrl = location.pathname;

  return (
    <Stack bg={"white"} _dark={{ bg: "gray.700" }} p={4} borderRadius={"lg"}>
      <Stack mt={4}>
        {navbarList.map((item) => (
          <Link
            key={item.name}
            onClick={() =>
              item.name === "Profil"
                ? navigate(`/${item.path}`)
                : navigate(item.path)
            }
            _hover={{ textDecoration: "none", bg: "primary" }}
            borderRadius={"lg"}
            py={3}
            bg={currentUrl === item.path ? "primary" : "transparent"}
          >
            <HStack ml={4}>
              <Text
                color={currentUrl === item.path ? "white" : "black"}
                _dark={{ color: "white" }}
              >
                {item.icon}
              </Text>
              <Text
                color={currentUrl === item.path ? "white" : "black"}
                _dark={{ color: "white" }}
              >
                {item.name}
              </Text>
            </HStack>
          </Link>
        ))}
      </Stack>
    </Stack>
  );
}
