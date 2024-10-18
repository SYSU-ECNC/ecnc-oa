import type { Metadata } from "next";
import "./globals.css";
import { ReactNode } from "react";

export const metadata: Metadata = {
  title: "ECNC OA",
  description: "用于 ECNC 排班，考勤，请假的系统",
};

const RootLayout = ({ children }: { children: ReactNode }) => {
  return (
    <html lang="zh-CN">
      <body>{children}</body>
    </html>
  );
};

export default RootLayout;
