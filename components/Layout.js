import Head from "next/head";
import styled from "styled-components";
import { useEffect } from "react";

export default function Layout({ children }) {
  return (
    <>
      <Head>
        <title>Alis Learning</title>
        <meta
          s
          name="description"
          content="A learning app to help children with dyslexia with their reading and writing"
        />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <header>
        <StyledTitle>Alis Learning</StyledTitle>
      </header>
      <main>{children}</main>
    </>
  );
}

const StyledTitle = styled.h1`
  color: black;
  font-family: OpenDyslexic;
  text-align: center;
`;
