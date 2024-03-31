import { Flex, Grid } from "@chakra-ui/react";

export default function Timeline({ children }: { children: React.ReactNode }) {
  return (
    <Grid gridArea="1/2/6/5">
      <Flex direction="column" grow={0} overflow="hidden" gap={4}>
        {children}
      </Flex>
    </Grid>
  );
}
