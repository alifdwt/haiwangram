import { ChakraProvider, extendTheme } from "@chakra-ui/react";

const chakraTheme = extendTheme({
  colors: {
    primary: "#1b79e5",
  },
});

export default function Providers({ children }: { children: React.ReactNode }) {
  return <ChakraProvider theme={chakraTheme}>{children}</ChakraProvider>;
}
