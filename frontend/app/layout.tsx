import "./styles/globals.css";
import Link from "next/link";
import GeoUpdater from "./GeoUpdater";

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
        <GeoUpdater />
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
