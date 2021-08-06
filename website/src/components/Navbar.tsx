import React from "react";
import NavbarBoot from "react-bootstrap/Navbar";
import Container from "react-bootstrap/Container";

export const Navbar: React.FC = () => {
  return (
    <NavbarBoot bg="dark" variant="dark">
      <Container>
        <NavbarBoot.Brand href="/">Web Scraper</NavbarBoot.Brand>
      </Container>
    </NavbarBoot>
  );
};
