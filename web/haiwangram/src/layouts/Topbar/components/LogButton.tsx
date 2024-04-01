import useToast from "@/hooks/useToast";
import { logout } from "@/slices/user/userSlice";
import { Button } from "@chakra-ui/react";
import { LogOutIcon } from "lucide-react";
import { useDispatch } from "react-redux";

function LogOutButton() {
  const dispatch = useDispatch();
  const toast = useToast();

  const handleClick = () => {
    dispatch(logout());
    toast("Success", "Logout berhasil, selamat tinggal!", "success");
  };

  return (
    <Button
      leftIcon={<LogOutIcon />}
      onClick={handleClick}
      colorScheme="red"
      w={"full"}
    >
      Keluar
    </Button>
  );
}

export { LogOutButton };
