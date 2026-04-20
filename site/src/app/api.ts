/**
 * API module for prompts and evaluations
 * Provides functions for fetching prompt and evaluation data
 */

export interface Prompt {
  id: string;
  name: string;
  service: string;
  language: string;
  plane: string;
  category?: string;
  tags?: string[];
  content?: string;
  difficulty?: string;
}

export interface Evaluation {
  promptId: string;
  configName: string;
  score: number;
  success?: boolean;
  timestamp?: string;
}

/**
 * Fetch all prompts
 */
export async function getPrompts(): Promise<Prompt[]> {
  const response = await fetch("/api/prompts");
  if (!response.ok) {
    throw new Error(`Failed to fetch prompts: ${response.statusText}`);
  }
  return response.json();
}

/**
 * Fetch all evaluations
 */
export async function getEvaluations(): Promise<Evaluation[]> {
  const response = await fetch("/api/evaluations");
  if (!response.ok) {
    throw new Error(`Failed to fetch evaluations: ${response.statusText}`);
  }
  return response.json();
}

/**
 * Fetch a specific prompt by ID
 */
export async function getPromptById(id: string): Promise<Prompt | null> {
  const response = await fetch(`/api/prompts/${id}`);
  if (!response.ok) {
    if (response.status === 404) return null;
    throw new Error(`Failed to fetch prompt ${id}: ${response.statusText}`);
  }
  return response.json();
}

/**
 * Fetch evaluations for a specific prompt
 */
export async function getEvaluationsForPrompt(promptId: string): Promise<Evaluation[]> {
  const response = await fetch(`/api/evaluations?promptId=${encodeURIComponent(promptId)}`);
  if (!response.ok) {
    throw new Error(`Failed to fetch evaluations for prompt ${promptId}: ${response.statusText}`);
  }
  return response.json();
}
