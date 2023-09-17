import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import "./globals.css";
import type { Metadata } from "next";
import { Inter } from "next/font/google";
import { cookies } from "next/headers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const [user, _] = await getUser();
  console.log({ user });
  return (
    <html lang="en" className="dark">
      <body className={inter.className}>
        <header>
          <nav className="flex flex-row">
            <Avatar className="ml-auto w-16 h-16">
              <AvatarImage src={user?.image} />
              <AvatarFallback>CN</AvatarFallback>
            </Avatar>
          </nav>
        </header>
        {children}
      </body>
    </html>
  );
}

interface User {
  name: string;
  email: string;
  image: string;
}
async function getUser() {
  const cookieStore = cookies();
  const token = cookieStore.get("oauth.session-token");
  try {
    const res = await fetch("http://localhost:8000/api/sessions/user", {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token?.value}`,
      },
    });
    const body = await res.json();
    console.log({ body });
    return [body, null] as [User, null];
  } catch (error) {
    console.log({ error });
    return [null, error] as [null, Error];
  }
}
