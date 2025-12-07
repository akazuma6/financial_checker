'use client';

import { useState } from 'react';
import {
  Container,
  Box,
  TextField,
  Button,
  Typography,
  Paper,
} from '@mui/material';
import { useRouter } from 'next/navigation';
import { useSnackbar } from 'notistack';

export default function Home() {
  const [code, setCode] = useState('');
  const router = useRouter();
  const { enqueueSnackbar } = useSnackbar();

  const handleSearch = () => {
    if (!code.trim()) {
      enqueueSnackbar('証券コードを入力してください', { variant: 'warning' });
      return;
    }
    router.push(`/companies/${code}`);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      handleSearch();
    }
  };

  return (
    <Container maxWidth="md" sx={{ mt: 8 }}>
      <Paper elevation={3} sx={{ p: 4 }}>
        <Box sx={{ textAlign: 'center', mb: 4 }}>
          <Typography variant="h3" component="h1" gutterBottom>
            Financial Checker
          </Typography>
          <Typography variant="h6" color="text.secondary">
            日本株財務健全性可視化プラットフォーム
          </Typography>
        </Box>

        <Box sx={{ mb: 4 }}>
          <Typography variant="body1" paragraph>
            証券コードまたは企業名で企業を検索し、財務データと健全性スコアを確認できます。
          </Typography>
        </Box>

        <Box sx={{ display: 'flex', flexDirection: { xs: 'column', sm: 'row' }, gap: 2, alignItems: 'center' }}>
          <Box sx={{ flex: { xs: '1', sm: '2' }, width: '100%' }}>
            <TextField
              fullWidth
              label="証券コード（例: 7203）"
              variant="outlined"
              value={code}
              onChange={(e) => setCode(e.target.value)}
              onKeyPress={handleKeyPress}
              placeholder="4桁の証券コードを入力"
            />
          </Box>
          <Box sx={{ flex: { xs: '1', sm: '1' }, width: { xs: '100%', sm: 'auto' } }}>
            <Button
              fullWidth
              variant="contained"
              size="large"
              onClick={handleSearch}
              sx={{ height: '56px' }}
            >
              検索
            </Button>
          </Box>
        </Box>

        <Box sx={{ mt: 4 }}>
          <Typography variant="body2" color="text.secondary">
            サンプルコード: 7203（トヨタ自動車）、6758（ソニーグループ）、9984（ソフトバンクグループ）
          </Typography>
        </Box>
      </Paper>
    </Container>
  );
}
