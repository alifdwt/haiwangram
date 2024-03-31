import { Box, Flex, Link, Stack, Text } from "@chakra-ui/react";
import { RssIcon, SearchSlashIcon } from "lucide-react";
import RegisterForm from "./RegisterForm";

const HeroSection = () => {
  return (
    <Box as="section" bg="white" _dark={{ bg: "gray.900" }} h={"92vh"}>
      <Flex
        maxW="screen-xl"
        px="4"
        pt="20"
        pb="8"
        mx="auto"
        gridGap="8"
        flexDirection={{ base: "column", lg: "row" }}
      >
        <Box mr="auto" alignSelf="center" flex="1">
          <Text
            maxW="2xl"
            mb="4"
            fontSize={{ base: "4xl", md: "5xl", xl: "6xl" }}
            fontWeight="extrabold"
            lineHeight="none"
            _dark={{ color: "white" }}
          >
            Temukan aksi <br />
            hewan luar biasa.
          </Text>
          <Text
            maxW="2xl"
            mb={{ base: "6", lg: "8" }}
            fontSize={{ base: "lg", lg: "xl" }}
            fontWeight="light"
            color="gray.500"
            _dark={{ color: "gray.400" }}
          >
            HaiwanGram adalah website yang dibangun dengan menggunakan{" "}
            <Link
              href="https://tailwindcss.com"
              color="blue.500"
              _hover={{ textDecoration: "underline" }}
            >
              cinta dan kasih sayang
            </Link>{" "}
            dan ditujukan untuk menyediakan informasi tentang para pemelihara
            hewan{" "}
            <Link
              href="https://flowbite.com/docs/getting-started/introduction/"
              color="blue.500"
              _hover={{ textDecoration: "underline" }}
            >
              domestik
            </Link>{" "}
            dan{" "}
            <Link
              href="https://flowbite.com/blocks/"
              color="blue.500"
              _hover={{ textDecoration: "underline" }}
            >
              liar
            </Link>
            .
          </Text>
          <Stack direction={{ base: "column", sm: "row" }} spacing="4">
            <Link
              href="https://github.com/themesberg/landwind"
              target="_blank"
              rel="noopener noreferrer"
              display="inline-flex"
              alignItems="center"
              justifyContent="center"
              w="full"
              px="5"
              py="3"
              gap={2}
              fontSize="sm"
              fontWeight="medium"
              textAlign="center"
              color="gray.900"
              border="1px solid"
              borderColor="gray.200"
              rounded="lg"
              _hover={{ bg: "gray.100" }}
              _dark={{
                color: "white",
                borderColor: "gray.700",
                focusRingColor: "gray.800",
              }}
            >
              <SearchSlashIcon /> Tentang Kami
            </Link>
            <Link
              href="https://www.figma.com/community/file/1125744163617429490"
              target="_blank"
              rel="noopener noreferrer"
              display="inline-flex"
              alignItems="center"
              justifyContent="center"
              w="full"
              px="5"
              py="3"
              gap={2}
              fontSize="sm"
              fontWeight="medium"
              textAlign="center"
              color="gray.900"
              bg="white"
              border="1px solid"
              borderColor="gray.200"
              rounded="lg"
              _hover={{ bg: "gray.100", color: "blue.700" }}
              _dark={{
                borderColor: "gray.600",
                _hover: { bg: "gray.700", color: "white" },
                focusRingColor: "gray.700",
              }}
            >
              <RssIcon />
              Blog
            </Link>
          </Stack>
        </Box>
        <Box
          display={{ base: "none", lg: "flex" }}
          mt={{ base: "0", lg: "auto" }}
          flex={"1"}
        >
          <Box
            bg={"gray.100"}
            _dark={{ bg: "gray.800" }}
            rounded={"xl"}
            h={"70vh"}
            w={"full"}
          >
            <RegisterForm />
          </Box>
        </Box>
      </Flex>
    </Box>
  );
};

export default HeroSection;
