import './styles/globals.css';

export const metadata = {
  title: 'Potarin',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ja">
      <body>{children}</body>
    </html>
  );
}
