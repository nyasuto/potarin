import "./styles/globals.css";
import Link from "next/link";

export const metadata = {
  title: "Potarin",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="ja">
      <body>
        <header className="p-4 bg-gray-100 mb-4">
          <Link href="/" className="font-bold">
            Potarin
          </Link>
        </header>
        {children}
      </body>
    </html>
  );
}
