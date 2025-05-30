"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { Search, Clock, Trash2, Send, ArrowRight } from "lucide-react";

export default function Home() {
  const [payloadId, setPayloadId] = useState("");
  const router = useRouter();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (payloadId) {
      router.push(`/payload/${payloadId}`);
    }
  };

  return (
    <main className="min-h-screen bg-background">
      {/* Hero Section */}
      <div className="w-[90%] max-w-2xl mx-auto pt-16 pb-12">
        <div className="space-y-6">
          <div className="space-y-3">
            <h1 className="text-4xl md:text-5xl font-bold text-foreground leading-tight">
              Rest
              <span className="text-neutral-600 dark:text-neutral-400 block text-2xl md:text-3xl font-medium mt-1 tracking-tight">
                Post. Expire. Vanish.
              </span>
            </h1>
            <p className="text-base text-neutral-600 dark:text-neutral-400 max-w-lg leading-relaxed">
              Share sensitive data securely with automatic expiration. Your
              payloads disappear after viewing or time limit.
            </p>
          </div>

          <form onSubmit={handleSubmit} className="space-y-3">
            <div className="relative">
              <input
                type="text"
                value={payloadId}
                onChange={(e) => setPayloadId(e.target.value)}
                placeholder="Enter your payload ID..."
                className="w-full px-5 py-3 pl-12 text-base text-foreground bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-lg focus:outline-none focus:ring-2 focus:ring-neutral-400 dark:focus:ring-neutral-600 transition-all duration-200"
              />
              <Search className="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-neutral-400" />
            </div>
            <button
              type="submit"
              className="w-full px-5 py-3 bg-neutral-900 dark:bg-white text-white dark:text-neutral-900 rounded-lg font-medium text-base hover:opacity-90 transition-all duration-200 flex items-center justify-center gap-2 group"
            >
              View Payload
              <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />
            </button>
          </form>
        </div>
      </div>

      {/* Features Section */}
      <div className="w-[90%] max-w-4xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
          <div className="p-6 bg-white dark:bg-neutral-900 rounded-lg border border-neutral-200 dark:border-neutral-800 hover:border-neutral-400 dark:hover:border-neutral-600 transition-colors duration-200">
            <div className="p-2.5 bg-neutral-100 dark:bg-neutral-800 rounded-lg w-fit mb-3">
              <Send className="w-5 h-5 text-neutral-600 dark:text-neutral-400" />
            </div>
            <h2 className="text-lg font-semibold text-foreground mb-1.5">
              Post Securely
            </h2>
            <p className="text-sm text-neutral-600 dark:text-neutral-400 leading-relaxed">
              Share your sensitive data with secure protocols
            </p>
          </div>
          <div className="p-6 bg-white dark:bg-neutral-900 rounded-lg border border-neutral-200 dark:border-neutral-800 hover:border-neutral-400 dark:hover:border-neutral-600 transition-colors duration-200">
            <div className="p-2.5 bg-neutral-100 dark:bg-neutral-800 rounded-lg w-fit mb-3">
              <Clock className="w-5 h-5 text-neutral-600 dark:text-neutral-400" />
            </div>
            <h2 className="text-lg font-semibold text-foreground mb-1.5">
              Set Expiration
            </h2>
            <p className="text-sm text-neutral-600 dark:text-neutral-400 leading-relaxed">
              Control how long your payloads remain accessible with time limits
            </p>
          </div>
          <div className="p-6 bg-white dark:bg-neutral-900 rounded-lg border border-neutral-200 dark:border-neutral-800 hover:border-neutral-400 dark:hover:border-neutral-600 transition-colors duration-200">
            <div className="p-2.5 bg-neutral-100 dark:bg-neutral-800 rounded-lg w-fit mb-3">
              <Trash2 className="w-5 h-5 text-neutral-600 dark:text-neutral-400" />
            </div>
            <h2 className="text-lg font-semibold text-foreground mb-1.5">
              Auto Vanish
            </h2>
            <p className="text-sm text-neutral-600 dark:text-neutral-400 leading-relaxed">
              Your data automatically disappears after viewing or time
              expiration
            </p>
          </div>
        </div>
      </div>
    </main>
  );
}
