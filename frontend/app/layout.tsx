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
        <header className="header-glass sticky top-0 z-50">
          <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div className="flex justify-between items-center py-4">
              <Link href="/" className="flex items-center space-x-2">
                <div className="w-8 h-8 hero-gradient rounded-full flex items-center justify-center floating-animation">
                  <span className="text-white font-bold text-sm">P</span>
                </div>
                <span className="text-2xl font-bold gradient-text">
                  Potarin
                </span>
              </Link>
              <nav className="hidden md:flex items-center space-x-8">
                <Link href="/" className="nav-link text-gray-700 hover:text-blue-600 font-medium transition-colors">
                  ホーム
                </Link>
                <Link href="/suggestions" className="nav-link text-gray-700 hover:text-blue-600 font-medium transition-colors">
                  コース提案
                </Link>
                <button className="btn-primary text-sm px-6 py-2">
                  始める
                </button>
              </nav>
              <button className="md:hidden p-2 rounded-lg hover:bg-gray-100 transition-colors">
                <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
                </svg>
              </button>
            </div>
          </div>
        </header>
        {children}
      </body>
    </html>
  );
}
