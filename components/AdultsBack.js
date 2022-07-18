import AdultBack from "./Styled-Components/AdultBack";
import { Router, useRouter } from "next/router";
import { useState } from "react";

export default function AdultsBack({ page }) {
  const router = useRouter();

  const navigateBack = () => {
    if (page > 1) {
      router.replace(`/adult-onboarding${page - 1}`);
    } else {
      router.replace("/");
    }
  };

  return <AdultBack onClick={navigateBack}>Back</AdultBack>;
}
