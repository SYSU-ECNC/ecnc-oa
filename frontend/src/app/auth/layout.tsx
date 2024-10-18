import LoginForm from "@/components/auth/login/login-form";
import Image from "next/image";
import { ReactNode } from "react";

interface AuthLayoutProps {
  children: ReactNode;
}

const AuthLayout = ({ children }: AuthLayoutProps) => {
  return (
    <div className="h-full grid grid-cols-2">
      {/* login form or forget password form */}
      <div className="grid place-items-center">{children}</div>
      {/* auth-background */}
      <div className="relative">
        <Image
          src="/auth-background.png"
          alt="auth background"
          fill={true}
          className="object-cover"
        />
      </div>
    </div>
  );
};

export default AuthLayout;
