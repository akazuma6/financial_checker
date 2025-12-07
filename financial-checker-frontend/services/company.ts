import { baseUrl, fetchWithQuery } from './utils';

export interface Company {
  code: string;
  name: string;
  industry: string;
}

export interface FinancialStatement {
  id: number;
  companyCode: string;
  fiscalYear: number;
  sales: number | null;
  operatingIncome: number | null;
  netIncome: number | null;
  netAssets: number | null;
  totalAssets: number | null;
  cashEquivalents: number | null;
  isConsolidated: boolean;
}

export interface HealthScore {
  companyCode: string;
  score: number;
  grade: string;
  equityRatio: number;
  currentRatio: number;
  roe: number;
  comment: string;
}

interface GetFinancialsResponse {
  status: string;
  message: string;
  data: FinancialStatement[];
}

interface GetHealthResponse {
  status: string;
  message: string;
  data: HealthScore;
}

class CompanyAPI {
  private static instance: CompanyAPI;
  private constructor() {}
  public static getInstance(): CompanyAPI {
    if (!CompanyAPI.instance) {
      CompanyAPI.instance = new CompanyAPI();
    }
    return CompanyAPI.instance;
  }

  async getFinancials(code: string): Promise<FinancialStatement[]> {
    const url = `${baseUrl}/api/v1/companies/${code}/financials`;
    const res = await fetchWithQuery<GetFinancialsResponse>({
      url,
      method: 'GET',
    });
    return res.data;
  }

  async getHealth(code: string): Promise<HealthScore> {
    const url = `${baseUrl}/api/v1/companies/${code}/health`;
    const res = await fetchWithQuery<GetHealthResponse>({
      url,
      method: 'GET',
    });
    return res.data;
  }
}

export const companyAPI = CompanyAPI.getInstance();
