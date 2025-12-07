import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import './globals.css';
import { ThemeProvider } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme } from './theme';
import { SnackbarProvider } from 'notistack';
import ClientProviders from './client-providers';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Financial Checker - 日本株財務健全性可視化プラットフォーム',
  description: 'EDINETの財務データを可視化し、企業の財務健全性を確認',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body className={inter.className}>
        <ClientProviders>
          {children}
        </ClientProviders>
      </body>
    </html>
  );
}
