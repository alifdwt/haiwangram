import {
  BrowserRouter,
  Navigate,
  Route,
  Routes as RouterRoutes,
} from "react-router-dom";
import Layout from "../layouts";
import { Heading } from "@chakra-ui/react";
import Home from "@/pages/Home/Home";

export default function Routes() {
  const accessToken = true;

  return (
    <BrowserRouter>
      <RouterRoutes>
        <Route element={accessToken ? <Layout /> : <Navigate to="/login" />}>
          <Route index element={<Home />} />
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
