import {
  Button,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  useDisclosure,
} from "@chakra-ui/react";

export default function Register() {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Button
        onClick={onOpen}
        bg={"primary"}
        color={"white"}
        _hover={{ bg: "primary" }}
        w={"md"}
      >
        Buat akun
      </Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Buat akun</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Fugiat,
            nostrum? Delectus explicabo aliquid maiores ipsam autem ab,
            recusandae laborum ipsum facilis error cum quaerat dolore doloribus
            sapiente nemo consequuntur corrupti quod, dicta natus. Laborum error
            fugit voluptatem eos, illo expedita voluptates, harum autem
            veritatis rem, nostrum aut eligendi omnis facilis!
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={onClose}>
              Close
            </Button>
            <Button variant="ghost">Secondary Action</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
