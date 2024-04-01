import {
  BrowserRouter,
  Navigate,
  Route,
  Routes as RouterRoutes,
} from "react-router-dom";
import Layout from "../layouts";
import { Heading } from "@chakra-ui/react";
import Home from "@/pages/Home/Home";
import Landing from "@/pages/Landing/Landing";
import Search from "@/pages/Search/Search";
import Follow from "@/pages/Follow/Follow";
import { useSelector } from "react-redux";
import { RootState } from "@/store";

export default function Routes() {
  const { accessToken } = useSelector((state: RootState) => state.user);

  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={accessToken ? <Layout /> : <Landing />}>
          <Route index element={<Home />} />
          <Route path="/search" element={<Search />} />
          <Route path="/follow" element={<Follow />} />
        </Route>
        <Route
          path="/login"
          element={accessToken ? <Navigate to="/" /> : <Heading>Login</Heading>}
        />
        <Route
          path="register"
          element={
            accessToken ? <Navigate to="/" /> : <Heading>Register</Heading>
          }
        />
      </RouterRoutes>
    </BrowserRouter>
  );
}
