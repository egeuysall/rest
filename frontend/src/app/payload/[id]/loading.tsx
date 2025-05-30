export default function Loading() {
  return (
    <div className="min-h-screen bg-background flex items-center justify-center">
      <div className="w-full max-w-3xl mx-auto px-4 py-12">
        <div className="bg-white dark:bg-neutral-900 rounded-lg shadow-lg border border-neutral-200 dark:border-neutral-800 overflow-hidden">
          <div className="p-8 border-b border-neutral-200 dark:border-neutral-800">
            <div className="h-8 w-48 bg-neutral-200 dark:bg-neutral-800 rounded animate-pulse mx-auto mb-6" />
            <div className="flex justify-center gap-8">
              <div className="h-4 w-32 bg-neutral-200 dark:bg-neutral-800 rounded animate-pulse" />
              <div className="h-4 w-40 bg-neutral-200 dark:bg-neutral-800 rounded animate-pulse" />
            </div>
          </div>
          <div className="p-8">
            <div className="h-96 bg-neutral-200 dark:bg-neutral-800 rounded animate-pulse" />
          </div>
        </div>
      </div>
    </div>
  );
}
