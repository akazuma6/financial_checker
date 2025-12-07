'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import {
  Container,
  Box,
  Typography,
  Paper,
  Grid,
  CircularProgress,
  Card,
  CardContent,
  Chip,
  Button,
} from '@mui/material';
import { useSnackbar } from 'notistack';
import {
  LineChart,
  Line,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from 'recharts';
import { companyAPI, FinancialStatement, HealthScore } from '@/services/company';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';

export default function CompanyDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { enqueueSnackbar } = useSnackbar();
  const code = params.code as string;

  const [financials, setFinancials] = useState<FinancialStatement[]>([]);
  const [health, setHealth] = useState<HealthScore | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const [financialsData, healthData] = await Promise.all([
          companyAPI.getFinancials(code),
          companyAPI.getHealth(code),
        ]);
        setFinancials(financialsData);
        setHealth(healthData);
      } catch (error) {
        console.error('Error fetching data:', error);
        enqueueSnackbar('データの取得に失敗しました', { variant: 'error' });
      } finally {
        setLoading(false);
      }
    };

    if (code) {
      fetchData();
    }
  }, [code, enqueueSnackbar]);

  const formatCurrency = (value: number | null) => {
    if (value === null) return 'N/A';
    return `${(value / 1000000000).toFixed(1)}億円`;
  };

  const getGradeColor = (grade: string) => {
    switch (grade) {
      case 'S':
        return 'success';
      case 'A':
        return 'info';
      case 'B':
        return 'warning';
      case 'C':
        return 'error';
      case 'D':
        return 'error';
      default:
        return 'default';
    }
  };

  if (loading) {
    return (
      <Container maxWidth="lg" sx={{ mt: 4, textAlign: 'center' }}>
        <CircularProgress />
      </Container>
    );
  }

  const chartData = financials
    .map((f) => ({
      year: f.fiscalYear,
      売上高: f.sales ? f.sales / 1000000000 : 0,
      営業利益: f.operatingIncome ? f.operatingIncome / 1000000000 : 0,
      当期純利益: f.netIncome ? f.netIncome / 1000000000 : 0,
    }))
    .reverse();

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Button
        startIcon={<ArrowBackIcon />}
        onClick={() => router.push('/')}
        sx={{ mb: 2 }}
      >
        戻る
      </Button>

      <Grid container spacing={3}>
        {/* 健全性スコア */}
        {health && (
          <Grid item xs={12}>
            <Card>
              <CardContent>
                <Typography variant="h5" gutterBottom>
                  財務健全性スコア
                </Typography>
                <Box sx={{ display: 'flex', alignItems: 'center', gap: 2, mt: 2 }}>
                  <Chip
                    label={`${health.grade} (${health.score}点)`}
                    color={getGradeColor(health.grade) as any}
                    size="large"
                    sx={{ fontSize: '1.2rem', height: '40px' }}
                  />
                  <Typography variant="body1" color="text.secondary">
                    {health.comment}
                  </Typography>
                </Box>
                <Box sx={{ mt: 2 }}>
                  <Typography variant="body2" color="text.secondary">
                    自己資本比率: {health.equityRatio.toFixed(1)}%
                  </Typography>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        )}

        {/* 財務データグラフ */}
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              財務データ推移（過去5年）
            </Typography>
            <ResponsiveContainer width="100%" height={400}>
              <LineChart data={chartData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="year" />
                <YAxis
                  label={{ value: '金額（億円）', angle: -90, position: 'insideLeft' }}
                />
                <Tooltip formatter={(value: number) => `${value.toFixed(1)}億円`} />
                <Legend />
                <Line
                  type="monotone"
                  dataKey="売上高"
                  stroke="#8884d8"
                  strokeWidth={2}
                />
                <Line
                  type="monotone"
                  dataKey="営業利益"
                  stroke="#82ca9d"
                  strokeWidth={2}
                />
                <Line
                  type="monotone"
                  dataKey="当期純利益"
                  stroke="#ffc658"
                  strokeWidth={2}
                />
              </LineChart>
            </ResponsiveContainer>
          </Paper>
        </Grid>

        {/* 財務データテーブル */}
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              財務データ詳細
            </Typography>
            <Box sx={{ overflowX: 'auto' }}>
              <table style={{ width: '100%', borderCollapse: 'collapse' }}>
                <thead>
                  <tr style={{ borderBottom: '2px solid #ddd' }}>
                    <th style={{ padding: '12px', textAlign: 'left' }}>年度</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>売上高</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>営業利益</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>当期純利益</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>純資産</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>総資産</th>
                    <th style={{ padding: '12px', textAlign: 'right' }}>現預金</th>
                  </tr>
                </thead>
                <tbody>
                  {financials
                    .sort((a, b) => b.fiscalYear - a.fiscalYear)
                    .map((f) => (
                      <tr key={f.id} style={{ borderBottom: '1px solid #eee' }}>
                        <td style={{ padding: '12px' }}>{f.fiscalYear}</td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.sales)}
                        </td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.operatingIncome)}
                        </td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.netIncome)}
                        </td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.netAssets)}
                        </td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.totalAssets)}
                        </td>
                        <td style={{ padding: '12px', textAlign: 'right' }}>
                          {formatCurrency(f.cashEquivalents)}
                        </td>
                      </tr>
                    ))}
                </tbody>
              </table>
            </Box>
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
}
