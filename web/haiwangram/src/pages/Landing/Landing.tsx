import Navbar from "./components/Navbar";
import HeroSection from "./components/Hero";
import { Box } from "@chakra-ui/react";

export default function Landing() {
  return (
    <Box>
      <Navbar />
      <HeroSection />
    </Box>
  );
}
