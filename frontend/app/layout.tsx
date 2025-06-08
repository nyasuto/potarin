import "./styles/globals.css";

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
          <a href="/" className="font-bold">
            Potarin
          </a>
        </header>
        {children}
      </body>
    </html>
  );
}
