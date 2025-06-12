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
      <body className="bg-gray-50">
        <header className="p-4 bg-gradient-to-r from-sky-500 to-green-400 text-white mb-8">
          <Link href="/" className="font-bold text-xl">
            Potarin
          </Link>
        </header>
        {children}
      </body>
    </html>
  );
}
