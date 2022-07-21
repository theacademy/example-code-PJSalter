import LearnerOnboardingLayout from "../../components/LearnerOnboardingLayout";
import { useEffect } from "react";

export default function OnboardingDone({ setOnboarding, background }) {
  useEffect(() => {
    document.body.style.backgroundColor = background;
  }, []);

  setOnboarding(true);

  return (
    <LearnerOnboardingLayout
      ttsTitle="Well done! Now you're ready to play a game."
      stepNumber={5}
      previousStep="background-selection"
      nextStep="child-landing"
      completed="true"
    />
  );
}
