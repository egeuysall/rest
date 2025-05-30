"use client";

import React, { useState, useMemo, useEffect } from "react";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import {
  dracula,
  oneLight,
} from "react-syntax-highlighter/dist/esm/styles/prism";
import { FileCode, FileText, Check, Copy, Code } from "lucide-react";
import { cn } from "@/lib/utils";

const languageIconMap: Record<string, React.ReactNode> = {
  html: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  css: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  js: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  jsx: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  javascript: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  ts: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  tsx: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  typescript: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  py: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  python: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  java: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  json: <FileCode className="w-4 h-4 text-black dark:text-white" />,
  txt: <FileText className="w-4 h-4 text-black dark:text-white" />,
  text: <FileText className="w-4 h-4 text-black dark:text-white" />,
  plaintext: <FileText className="w-4 h-4 text-black dark:text-white" />,
  md: <FileText className="w-4 h-4 text-black dark:text-white" />,
  markdown: <FileText className="w-4 h-4 text-black dark:text-white" />,
  sh: <FileText className="w-4 h-4 text-black dark:text-white" />,
  bash: <FileText className="w-4 h-4 text-black dark:text-white" />,
  shell: <FileText className="w-4 h-4 text-black dark:text-white" />,
};

type CodeBlockProps = {
  code: string;
  language?: string;
  fileName?: string | null;
  className?: string;
  codeClassName?: string;
};

const CodeBlock: React.FC<CodeBlockProps> = ({
  code = "print('Hello world!')",
  language = "py",
  fileName = "hello",
  className,
  codeClassName,
}) => {
  const [copied, setCopied] = useState(false);
  const [isClient, setIsClient] = useState(false);
  const [isDarkMode, setIsDarkMode] = useState(false);

  useEffect(() => {
    setIsClient(true);
    setIsDarkMode(
      window.matchMedia &&
        window.matchMedia("(prefers-color-scheme: dark)").matches
    );
    const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");
    const listener = (e: MediaQueryListEvent) => setIsDarkMode(e.matches);
    mediaQuery.addEventListener("change", listener);
    return () => mediaQuery.removeEventListener("change", listener);
  }, []);

  const handleCopy = async () => {
    if (!isClient) return;
    try {
      if (navigator.clipboard && navigator.clipboard.writeText) {
        await navigator.clipboard.writeText(code);
      } else {
        const textArea = document.createElement("textarea");
        textArea.value = code;
        textArea.style.top = "0";
        textArea.style.left = "0";
        textArea.style.position = "fixed";
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        document.execCommand("copy");
        document.body.removeChild(textArea);
      }
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch (error) {
      console.error("Failed to copy code:", error);
    }
  };

  const langIcon = useMemo(() => {
    const lang = language?.toLowerCase();
    if (lang && languageIconMap[lang]) return languageIconMap[lang];
    const ext = lang?.split(/[^a-z]/gi)[0];
    if (ext && languageIconMap[ext]) return languageIconMap[ext];
    return <Code className="w-4 h-4 text-black dark:text-white" />;
  }, [language]);

  const prismTheme = isDarkMode ? dracula : oneLight;
  const codeLines = code.split("\n");
  const lineNumbers = Array.from({ length: codeLines.length }, (_, i) => i + 1);

  return (
    <figure className={cn("relative w-full", className)}>
      <div
        className={cn(
          "rounded-lg overflow-hidden w-full border border-neutral-200 dark:border-neutral-800 bg-transparent shadow-sm hover:shadow-md transition-shadow duration-200"
        )}
      >
        {fileName && (
          <figcaption
            className={cn(
              "flex items-center justify-between px-4 py-3 border-b border-neutral-200 dark:border-neutral-800 bg-transparent backdrop-blur-sm"
            )}
          >
            <div className="flex items-center gap-2">
              <div className="flex-shrink-0">{langIcon}</div>
              <span className="text-sm font-medium text-neutral-900 dark:text-neutral-100">
                {`${fileName}${language ? `.${language}` : ""}`}
              </span>
            </div>
            <button
              onClick={handleCopy}
              aria-label={copied ? "Copied!" : "Copy code"}
              className={cn(
                "flex items-center transition-all duration-200 hover:opacity-80 text-neutral-600 dark:text-neutral-400 ml-6 hover:scale-110"
              )}
              title={copied ? "Copied!" : "Copy code"}
              type="button"
              disabled={!isClient}
            >
              {copied ? (
                <Check className="w-4 h-4" />
              ) : (
                <Copy className="w-4 h-4" />
              )}
            </button>
          </figcaption>
        )}
        <div className={cn("relative group")}>
          <div className="max-h-[500px] overflow-y-auto scrollbar-thin scrollbar-thumb-neutral-300 dark:scrollbar-thumb-neutral-700 scrollbar-track-transparent hover:scrollbar-thumb-neutral-400 dark:hover:scrollbar-thumb-neutral-600">
            <div
              className={cn(
                "grid grid-cols-[auto_1fr] w-full font-mono",
                codeClassName
              )}
            >
              <div
                className={cn(
                  "flex flex-col items-end pt-4 pb-4 px-3 bg-transparent select-none text-neutral-500 dark:text-neutral-400 border-r border-neutral-200 dark:border-neutral-800 font-mono sticky left-0"
                )}
              >
                {lineNumbers.map((num) => (
                  <div
                    key={num}
                    className="leading-[1.5] text-sm min-h-[1.5em]"
                  >
                    {num}
                  </div>
                ))}
              </div>
              <div className="relative flex-1 min-w-0 pt-4 pb-4 pl-4 pr-4 overflow-x-auto">
                {isClient && (
                  <SyntaxHighlighter
                    language={language}
                    style={prismTheme}
                    customStyle={{
                      margin: 0,
                      padding: 0,
                      background: "transparent",
                      fontSize: "0.875rem",
                      fontFamily: "inherit",
                    }}
                    codeTagProps={{
                      style: {
                        fontFamily: "inherit",
                      },
                    }}
                    showLineNumbers={false}
                    wrapLongLines={false}
                    lineProps={() => ({
                      style: {
                        whiteSpace: "pre",
                        minHeight: "1.5em",
                        marginBottom: 0,
                      },
                    })}
                  >
                    {code}
                  </SyntaxHighlighter>
                )}
              </div>
            </div>
          </div>
          {!fileName && isClient && (
            <button
              onClick={handleCopy}
              aria-label={copied ? "Copied!" : "Copy code"}
              className={cn(
                "absolute right-2 top-2 transition-all duration-200 hover:opacity-80 text-neutral-600 dark:text-neutral-400 hover:scale-110"
              )}
              title={copied ? "Copied!" : "Copy code"}
              type="button"
              disabled={!isClient}
            >
              {copied ? (
                <Check className="w-5 h-5" />
              ) : (
                <Copy className="w-4 h-4" />
              )}
            </button>
          )}
        </div>
      </div>
    </figure>
  );
};

export default CodeBlock;
