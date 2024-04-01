import { ChakraProvider, extendTheme } from "@chakra-ui/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider } from "react-redux";
import { PersistGate } from "redux-persist/integration/react";
import { persistor, store } from "./store";

const chakraTheme = extendTheme({
  colors: {
    primary: {
      "50": "#eff9ff",
      "100": "#daf2ff",
      "200": "#bee8ff",
      "300": "#91dbff",
      "400": "#5dc6fd",
      "500": "#37a9fa",
      "600": "#218bef",
      "700": "#1b79e5",
      "800": "#1b5db2",
      "900": "#1c508c",
      "950": "#163155",
    },
  },
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      refetchOnWindowFocus: false,
    },
  },
});

export default function Providers({ children }: { children: React.ReactNode }) {
  return (
    <Provider store={store}>
      <PersistGate loading={null} persistor={persistor}>
        <QueryClientProvider client={queryClient}>
          <ChakraProvider theme={chakraTheme}>{children}</ChakraProvider>
        </QueryClientProvider>
      </PersistGate>
    </Provider>
  );
}
