import type { Metadata } from "next";
import "./globals.css";
import { ReactNode } from "react";
import localFont from "next/font/local";

const LXGWWenKai = localFont({
  src: [
    {
      path: "./fonts/LXGWWenKai/LXGWWenKai-Light.ttf",
      weight: "300",
      style: "normal",
    },
    {
      path: "./fonts/LXGWWenKai/LXGWWenKai-Regular.ttf",
      weight: "400",
      style: "normal",
    },
    {
      path: "./fonts/LXGWWenKai/LXGWWenKai-Medium.ttf",
      weight: "500",
      style: "normal",
    },
  ],
});

export const metadata: Metadata = {
  title: "ECNC OA",
  description: "用于 ECNC 排班，考勤，请假的系统",
};

interface RootLayoutProps {
  children: ReactNode;
}

const RootLayout = ({ children }: RootLayoutProps) => {
  return (
    <html lang="zh-CN">
      <body className={`${LXGWWenKai.className} antialiased`}>{children}</body>
    </html>
  );
};

export default RootLayout;
