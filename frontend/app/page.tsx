import { Suggestion } from "potarin-shared/types";
import Link from "next/link";
import { ensureUniqueIds } from "../lib/ensureUniqueIds";

async function getSuggestions(): Promise<Suggestion[]> {
  const res = await fetch(
    `${process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080"}/api/v1/suggestions`,
    { cache: "no-store" },
  );
  if (!res.ok) {
    throw new Error("Failed to fetch suggestions");
  }
  const data = await res.json();
  return Array.isArray(data) ? data : [];
}

export default async function Home() {
  const suggestions = ensureUniqueIds(await getSuggestions());

  return (
    <div className="min-h-screen hero-background">
      {/* Hero Section */}
      <section className="relative overflow-hidden">
        <div className="relative max-w-7xl mx-auto px-4 py-20 sm:px-6 lg:px-8">
          <div className="text-center">
            <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6 floating-animation">
              <span className="gradient-text">AI</span>が提案する
              <br />あなただけの
              <span className="gradient-text">サイクリング</span>コース
            </h1>
            <p className="text-xl text-gray-600 mb-8 max-w-3xl mx-auto leading-relaxed">
              希望や現在地から、AIが最適なサイクリングルートを提案。
              地図上でルートを確認して、新しい発見の旅に出かけましょう。
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
              <button className="btn-primary">
                🚴‍♂️ コース提案を受ける
              </button>
              <button className="btn-secondary">
                📍 現在地から探す
              </button>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-white">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">Potarinでできること</h2>
            <p className="text-lg text-gray-600">AIの力で、あなたにぴったりのサイクリング体験を</p>
          </div>
          <div className="grid md:grid-cols-3 gap-8">
            <div className="text-center p-6 rounded-xl bg-blue-50 feature-card">
              <div className="w-16 h-16 bg-blue-600 rounded-full flex items-center justify-center mx-auto mb-4 floating-animation">
                <span className="text-2xl text-white">🤖</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">AI提案</h3>
              <p className="text-gray-600">あなたの好みや体力レベルに合わせて、AIが最適なコースを提案します</p>
            </div>
            <div className="text-center p-6 rounded-xl bg-green-50 feature-card">
              <div className="w-16 h-16 bg-green-600 rounded-full flex items-center justify-center mx-auto mb-4 floating-animation" style={{animationDelay: "0.5s"}}>
                <span className="text-2xl text-white">🗺️</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">地図表示</h3>
              <p className="text-gray-600">提案されたルートを地図上で確認。詳細な道順も一目で分かります</p>
            </div>
            <div className="text-center p-6 rounded-xl bg-purple-50 feature-card">
              <div className="w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center mx-auto mb-4 floating-animation" style={{animationDelay: "1s"}}>
                <span className="text-2xl text-white">📍</span>
              </div>
              <h3 className="text-xl font-semibold text-gray-900 mb-2">現在地連携</h3>
              <p className="text-gray-600">あなたの現在地から出発できるコースを優先的に提案します</p>
            </div>
          </div>
        </div>
      </section>

      {/* Course Suggestions Section */}
      <section className="py-16 bg-gray-50">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="text-center mb-12">
            <h2 className="text-3xl font-bold text-gray-900 mb-4">おすすめコース</h2>
            <p className="text-lg text-gray-600">AIが今日のあなたにおすすめするサイクリングコース</p>
          </div>
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {suggestions.map((s) => (
              <Link key={s.id} href={`/suggestions/${s.id}`}>
                <div className="course-card">
                  <div className="course-image">
                    <div className="absolute bottom-4 left-4 text-white z-10">
                      <span className="inline-block bg-white/20 px-3 py-1 rounded-full text-sm font-medium backdrop-blur-sm">
                        🚴‍♂️ サイクリング
                      </span>
                    </div>
                  </div>
                  <div className="p-6">
                    <h3 className="text-xl font-semibold text-gray-900 mb-2 group-hover:text-blue-600 transition-colors">
                      {s.title}
                    </h3>
                    <p className="text-gray-600 leading-relaxed">{s.description}</p>
                    <div className="mt-4 flex items-center text-blue-600 font-medium">
                      <span>詳細を見る</span>
                      <svg className="w-4 h-4 ml-2 group-hover:translate-x-1 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 5l7 7-7 7" />
                      </svg>
                    </div>
                  </div>
                </div>
              </Link>
            ))}
          </div>
        </div>
      </section>
    </div>
  );
}
