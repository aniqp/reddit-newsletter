import React from "react";
import LoginCard from "@/components/LoginCard";
import AppDescription from "@/components/AppDescription";

export default function Login() {
  return (
    <div className="w-full h-screen flex items-center">
      <div className="h-full w-1/2">
        <AppDescription />
      </div>
      <div className="w-1/2">
        <LoginCard />
      </div>
    </div>
  );
}
