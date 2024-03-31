import { Box, Button, Flex, Link, Text, useColorMode } from "@chakra-ui/react";
import { GithubIcon, MoonIcon, PawPrintIcon, SunIcon } from "lucide-react";

const navbarList = [
  {
    name: "Beranda",
    href: "/",
  },
  {
    name: "Tentang Kami",
    href: "/about",
  },
  {
    name: "Blog",
    href: "/blog",
  },
  {
    name: "FAQ",
    href: "/faq",
  },
  {
    name: "Toko",
    href: "/shop",
  },
];

export default function Navbar() {
  const { colorMode, toggleColorMode } = useColorMode();
  const currentUrl = location.pathname;

  return (
    <Box width={"full"}>
      <Box bg={"white"} border={"gray.200"} py={2} _dark={{ bg: "gray.900" }}>
        <Flex
          wrap={"wrap"}
          alignItems={"center"}
          justifyContent={"space-between"}
          maxW={"6xl"}
          mx={"auto"}
          px={4}
        >
          <Link href="/" as={Flex} alignItems={"center"} gap={2}>
            <Box bg={"primary"} borderRadius={"lg"} color={"white"} p={2}>
              <PawPrintIcon />
            </Box>
            <Text fontSize="2xl" fontWeight="bold">
              HaiwanGram
            </Text>
          </Link>
          <Flex alignItems={"center"} order={{ lg: 2 }} gap={2}>
            {/* <Box display={{ base: "none", md: "block" }}> */}
            <Button onClick={toggleColorMode}>
              {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
            </Button>
            <Button leftIcon={<GithubIcon />} variant="outline">
              alifdwt
            </Button>
            {/* </Box> */}
            <Button bg={"primary"} color={"white"} _hover={{ bg: "blue.600" }}>
              Register
            </Button>
          </Flex>
          <Box
            display={{ base: "hidden", lg: "flex" }}
            alignItems={"center"}
            justifyContent={"space-between"}
            w={{ base: "full", lg: "auto" }}
            order={{ lg: 1 }}
          >
            <Flex
              flexDir={{ base: "column", lg: "row" }}
              mt={{ base: 4, lg: 0 }}
            >
              {navbarList.map((item) => (
                <Link
                  href={item.href}
                  key={item.name}
                  _hover={{ textDecoration: "none" }}
                >
                  <Text
                    fontSize="lg"
                    fontWeight="medium"
                    color={currentUrl === item.href ? "primary" : ""}
                    _hover={{ color: "primary" }}
                    py={2}
                    pl={3}
                    pr={4}
                  >
                    {item.name}
                  </Text>
                </Link>
              ))}
            </Flex>
          </Box>
        </Flex>
      </Box>
    </Box>
  );
}
