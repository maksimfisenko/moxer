import { Link } from "@chakra-ui/react";
import { LuGithub } from "react-icons/lu";
import { SocialIcon } from "react-social-icons";

const Socials = () => {
  return (
    <>
      <Link href="https://github.com/maksimfisenko/moxer" target="_blank">
        <LuGithub />
      </Link>
      <SocialIcon
        target="_blank"
        url="https://t.me/maximfisenko"
        style={{ height: "20px", width: "20px" }}
      />
    </>
  );
};

export default Socials;
