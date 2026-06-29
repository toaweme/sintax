package text

import (
	"testing"

	"github.com/toaweme/sintax/assert"
)

func Test_ModelTitle_EdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "O1-2025-12-11", expected: "o1 (2025-12-11)"},
		{input: "o1-2025-12-11", expected: "o1 (2025-12-11)"},
		{input: "o1-2025-12-11", expected: "o1 (2025-12-11)"},
		{input: "gpt-5-2025-08-07", expected: "GPT 5 (2025-08-07)"},
		{input: "gpt-5.2-2025-12-11", expected: "GPT 5.2 (2025-12-11)"},
		{input: "gpt-5.2-20251211", expected: "GPT 5.2 (2025-12-11)"},
		{input: "openai/chatgpt-4o-latest", expected: "OpenAI: ChatGPT 4o Latest"},
		{input: "openai/gpt-5-2025-08-07", expected: "OpenAI: GPT 5 (2025-08-07)"},
		{input: "openai/gpt-o5-20250807", expected: "OpenAI: GPT o5 (2025-08-07)"},
		{input: "openai/gpt-5-2025-08-07:free", expected: "OpenAI: GPT 5 (2025-08-07) (free)"},
		{input: "openai/o1-2025-12-11", expected: "OpenAI: o1 (2025-12-11)"},

		// some randoms to test we didn't break anything
		{input: "allenai/olmo-2-0325-32b-instruct", expected: "Allenai: Olmo 2-0325 32B Instruct"},
		{input: "anthracite-org/magnum-v4-72b", expected: "Anthracite Org: Magnum v4 72B"},
		{input: "deepseek/deepseek-r1-0528-qwen3-8b", expected: "Deepseek: Deepseek R1 0528 Qwen3 8B"},
		{input: "qwen/qwen-2.5-72b-instruct", expected: "Qwen: Qwen 2.5 72B Instruct"},
		{input: "qwen/qwen-2.5-7b-instruct", expected: "Qwen: Qwen 2.5 7B Instruct"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatModelTitle(tt.input))
		})
	}
}

func Test_ModelTitle_RealSlugData(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "ai21/jamba-large-1.7", expected: "AI21: Jamba Large 1.7"},
		{input: "ai21/jamba-mini-1.7", expected: "AI21: Jamba Mini 1.7"},
		{input: "aion-labs/aion-1.0", expected: "AionLabs: Aion 1.0"},
		{input: "aion-labs/aion-1.0-mini", expected: "AionLabs: Aion 1.0 Mini"},
		{input: "aion-labs/aion-rp-llama-3.1-8b", expected: "AionLabs: Aion Rp Llama 3.1 8B"},
		{input: "alfredpros/codellama-7b-instruct-solidity", expected: "Alfredpros: Codellama 7B Instruct Solidity"},
		{input: "alibaba/tongyi-deepresearch-30b-a3b", expected: "Alibaba: Tongyi Deepresearch 30B A3B"},
		{input: "alibaba/tongyi-deepresearch-30b-a3b:free", expected: "Alibaba: Tongyi Deepresearch 30B A3B (free)"},
		{input: "allenai/olmo-2-0325-32b-instruct", expected: "Allenai: Olmo 2-0325 32B Instruct"},
		{input: "allenai/olmo-3-32b-think:free", expected: "Allenai: Olmo 3 32B Think (free)"},
		{input: "allenai/olmo-3-7b-instruct", expected: "Allenai: Olmo 3 7B Instruct"},
		{input: "allenai/olmo-3-7b-think", expected: "Allenai: Olmo 3 7B Think"},
		{input: "allenai/olmo-3.1-32b-think:free", expected: "Allenai: Olmo 3.1 32B Think (free)"},
		{input: "alpindale/goliath-120b", expected: "Alpindale: Goliath 120B"},
		{input: "amazon/nova-2-lite-v1", expected: "Amazon: Nova 2 Lite v1"},
		{input: "amazon/nova-lite-v1", expected: "Amazon: Nova Lite v1"},
		{input: "amazon/nova-micro-v1", expected: "Amazon: Nova Micro v1"},
		{input: "amazon/nova-premier-v1", expected: "Amazon: Nova Premier v1"},
		{input: "amazon/nova-pro-v1", expected: "Amazon: Nova Pro v1"},
		{input: "anthracite-org/magnum-v4-72b", expected: "Anthracite Org: Magnum v4 72B"},
		{input: "anthropic/claude-3-haiku", expected: "Anthropic: Claude 3 Haiku"},
		{input: "anthropic/claude-3-opus", expected: "Anthropic: Claude 3 Opus"},
		{input: "anthropic/claude-3.5-haiku", expected: "Anthropic: Claude 3.5 Haiku"},
		{input: "anthropic/claude-3.5-haiku-20241022", expected: "Anthropic: Claude 3.5 Haiku (2024-10-22)"},
		{input: "anthropic/claude-3.5-sonnet", expected: "Anthropic: Claude 3.5 Sonnet"},
		{input: "anthropic/claude-3.7-sonnet", expected: "Anthropic: Claude 3.7 Sonnet"},
		{input: "anthropic/claude-3.7-sonnet:thinking", expected: "Anthropic: Claude 3.7 Sonnet Thinking"},
		{input: "anthropic/claude-haiku-4.5", expected: "Anthropic: Claude Haiku 4.5"},
		{input: "anthropic/claude-opus-4", expected: "Anthropic: Claude Opus 4"},
		{input: "anthropic/claude-opus-4.1", expected: "Anthropic: Claude Opus 4.1"},
		{input: "anthropic/claude-opus-4.5", expected: "Anthropic: Claude Opus 4.5"},
		{input: "anthropic/claude-sonnet-4", expected: "Anthropic: Claude Sonnet 4"},
		{input: "anthropic/claude-sonnet-4.5", expected: "Anthropic: Claude Sonnet 4.5"},
		{input: "arcee-ai/coder-large", expected: "ArceeAI: Coder Large"},
		{input: "arcee-ai/maestro-reasoning", expected: "ArceeAI: Maestro Reasoning"},
		{input: "arcee-ai/spotlight", expected: "ArceeAI: Spotlight"},
		{input: "arcee-ai/trinity-mini", expected: "ArceeAI: Trinity Mini"},
		{input: "arcee-ai/trinity-mini:free", expected: "ArceeAI: Trinity Mini (free)"},
		{input: "arcee-ai/virtuoso-large", expected: "ArceeAI: Virtuoso Large"},
		{input: "arliai/qwq-32b-arliai-rpr-v1", expected: "Arliai: Qwq 32B Arliai Rpr v1"},
		{input: "baidu/ernie-4.5-21b-a3b", expected: "Baidu: Ernie 4.5 21B A3B"},
		{input: "baidu/ernie-4.5-21b-a3b-thinking", expected: "Baidu: Ernie 4.5 21B A3B Thinking"},
		{input: "baidu/ernie-4.5-300b-a47b", expected: "Baidu: Ernie 4.5 300B A47B"},
		{input: "baidu/ernie-4.5-vl-28b-a3b", expected: "Baidu: Ernie 4.5 Vl 28B A3B"},
		{input: "baidu/ernie-4.5-vl-424b-a47b", expected: "Baidu: Ernie 4.5 Vl 424B A47B"},
		{input: "bytedance/ui-tars-1.5-7b", expected: "ByteDance: Ui TARS 1.5 7B"},
		{input: "cognitivecomputations/dolphin-mistral-24b-venice-edition:free", expected: "Cognitivecomputations: Dolphin Mistral 24B Venice Edition (free)"},
		{input: "cohere/command-a", expected: "Cohere: Command A"},
		{input: "cohere/command-r-08-2024", expected: "Cohere: Command R (2024-08)"},
		{input: "cohere/command-r-plus-08-2024", expected: "Cohere: Command R Plus (2024-08)"},
		{input: "cohere/command-r7b-12-2024", expected: "Cohere: Command R7B (2024-12)"},
		{input: "deepcogito/cogito-v2-preview-llama-109b-moe", expected: "Deepcogito: Cogito v2 Preview Llama 109B Moe"},
		{input: "deepcogito/cogito-v2-preview-llama-405b", expected: "Deepcogito: Cogito v2 Preview Llama 405B"},
		{input: "deepcogito/cogito-v2-preview-llama-70b", expected: "Deepcogito: Cogito v2 Preview Llama 70B"},
		{input: "deepcogito/cogito-v2.1-671b", expected: "Deepcogito: Cogito v2.1 671B"},
		{input: "deepseek/deepseek-chat", expected: "Deepseek: Deepseek Chat"},
		{input: "deepseek/deepseek-chat-v3-0324", expected: "Deepseek: Deepseek Chat v3 0324"},
		{input: "deepseek/deepseek-chat-v3.1", expected: "Deepseek: Deepseek Chat v3.1"},
		{input: "deepseek/deepseek-prover-v2", expected: "Deepseek: Deepseek Prover v2"},
		{input: "deepseek/deepseek-r1", expected: "Deepseek: Deepseek R1"},
		{input: "deepseek/deepseek-r1-0528", expected: "Deepseek: Deepseek R1 0528"},
		{input: "deepseek/deepseek-r1-0528-qwen3-8b", expected: "Deepseek: Deepseek R1 0528 Qwen3 8B"},
		{input: "deepseek/deepseek-r1-0528:free", expected: "Deepseek: Deepseek R1 0528 (free)"},
		{input: "deepseek/deepseek-r1-distill-llama-70b", expected: "Deepseek: Deepseek R1 Distill Llama 70B"},
		{input: "deepseek/deepseek-r1-distill-qwen-14b", expected: "Deepseek: Deepseek R1 Distill Qwen 14B"},
		{input: "deepseek/deepseek-r1-distill-qwen-32b", expected: "Deepseek: Deepseek R1 Distill Qwen 32B"},
		{input: "deepseek/deepseek-v3.1-terminus", expected: "Deepseek: Deepseek v3.1 Terminus"},
		{input: "deepseek/deepseek-v3.1-terminus:exacto", expected: "Deepseek: Deepseek v3.1 Terminus (exacto)"},
		{input: "deepseek/deepseek-v3.2", expected: "Deepseek: Deepseek v3.2"},
		{input: "deepseek/deepseek-v3.2-exp", expected: "Deepseek: Deepseek v3.2 Exp"},
		{input: "deepseek/deepseek-v3.2-speciale", expected: "Deepseek: Deepseek v3.2 Speciale"},
		{input: "eleutherai/llemma_7b", expected: "EleutherAI: Llemma 7B"},
		{input: "essentialai/rnj-1-instruct", expected: "EssentialAI: Rnj 1 Instruct"},
		{input: "google/gemini-2.0-flash-001", expected: "Google: Gemini 2.0 Flash 001"},
		{input: "google/gemini-2.0-flash-exp:free", expected: "Google: Gemini 2.0 Flash Exp (free)"},
		{input: "google/gemini-2.0-flash-lite-001", expected: "Google: Gemini 2.0 Flash Lite 001"},
		{input: "google/gemini-2.5-flash", expected: "Google: Gemini 2.5 Flash"},
		{input: "google/gemini-2.5-flash-image", expected: "Google: Gemini 2.5 Flash Image"},
		{input: "google/gemini-2.5-flash-image-preview", expected: "Google: Gemini 2.5 Flash Image Preview"},
		{input: "google/gemini-2.5-flash-lite", expected: "Google: Gemini 2.5 Flash Lite"},
		{input: "google/gemini-2.5-flash-lite-preview-09-2025", expected: "Google: Gemini 2.5 Flash Lite Preview (2025-09)"},
		{input: "google/gemini-2.5-flash-preview-09-2025", expected: "Google: Gemini 2.5 Flash Preview (2025-09)"},
		{input: "google/gemini-2.5-pro", expected: "Google: Gemini 2.5 Pro"},
		{input: "google/gemini-2.5-pro-preview", expected: "Google: Gemini 2.5 Pro Preview"},
		{input: "google/gemini-2.5-pro-preview-05-06", expected: "Google: Gemini 2.5 Pro Preview 05-06"},
		{input: "google/gemini-3-flash-preview", expected: "Google: Gemini 3 Flash Preview"},
		{input: "google/gemini-3-pro-image-preview", expected: "Google: Gemini 3 Pro Image Preview"},
		{input: "google/gemini-3-pro-preview", expected: "Google: Gemini 3 Pro Preview"},
		{input: "google/gemma-2-27b-it", expected: "Google: Gemma 2 27B Instruct"},
		{input: "google/gemma-2-9b-it", expected: "Google: Gemma 2 9B Instruct"},
		{input: "google/gemma-3-12b-it", expected: "Google: Gemma 3 12B Instruct"},
		{input: "google/gemma-3-12b-it:free", expected: "Google: Gemma 3 12B Instruct (free)"},
		{input: "google/gemma-3-27b-it", expected: "Google: Gemma 3 27B Instruct"},
		{input: "google/gemma-3-27b-it:free", expected: "Google: Gemma 3 27B Instruct (free)"},
		{input: "google/gemma-3-4b-it", expected: "Google: Gemma 3 4B Instruct"},
		{input: "google/gemma-3-4b-it:free", expected: "Google: Gemma 3 4B Instruct (free)"},
		{input: "google/gemma-3n-e2b-it:free", expected: "Google: Gemma 3n E2B Instruct (free)"},
		{input: "google/gemma-3n-e4b-it", expected: "Google: Gemma 3n E4B Instruct"},
		{input: "google/gemma-3n-e4b-it:free", expected: "Google: Gemma 3n E4B Instruct (free)"},
		{input: "gryphe/mythomax-l2-13b", expected: "Gryphe: Mythomax L2 13B"},
		{input: "ibm-granite/granite-4.0-h-micro", expected: "Ibm Granite: Granite 4.0 H Micro"},
		{input: "inception/mercury", expected: "Inception: Mercury"},
		{input: "inception/mercury-coder", expected: "Inception: Mercury Coder"},
		{input: "inflection/inflection-3-pi", expected: "Inflection: Inflection 3 Pi"},
		{input: "inflection/inflection-3-productivity", expected: "Inflection: Inflection 3 Productivity"},
		{input: "kwaipilot/kat-coder-pro:free", expected: "Kwaipilot: KAT Coder Pro (free)"},
		{input: "liquid/lfm-2.2-6b", expected: "Liquid: LFM 2.2 6B"},
		{input: "liquid/lfm2-8b-a1b", expected: "Liquid: LFM2 8B A1B"},
		{input: "mancer/weaver", expected: "Mancer: Weaver"},
		{input: "meituan/longcat-flash-chat", expected: "Meituan: Longcat Flash Chat"},
		{input: "meta-llama/llama-3-70b-instruct", expected: "Meta Llama: Llama 3 70B Instruct"},
		{input: "meta-llama/llama-3-8b-instruct", expected: "Meta Llama: Llama 3 8B Instruct"},
		{input: "meta-llama/llama-3.1-405b", expected: "Meta Llama: Llama 3.1 405B"},
		{input: "meta-llama/llama-3.1-405b-instruct", expected: "Meta Llama: Llama 3.1 405B Instruct"},
		{input: "meta-llama/llama-3.1-405b-instruct:free", expected: "Meta Llama: Llama 3.1 405B Instruct (free)"},
		{input: "meta-llama/llama-3.1-70b-instruct", expected: "Meta Llama: Llama 3.1 70B Instruct"},
		{input: "meta-llama/llama-3.1-8b-instruct", expected: "Meta Llama: Llama 3.1 8B Instruct"},
		{input: "meta-llama/llama-3.2-11b-vision-instruct", expected: "Meta Llama: Llama 3.2 11B Vision Instruct"},
		{input: "meta-llama/llama-3.2-1b-instruct", expected: "Meta Llama: Llama 3.2 1B Instruct"},
		{input: "meta-llama/llama-3.2-3b-instruct", expected: "Meta Llama: Llama 3.2 3B Instruct"},
		{input: "meta-llama/llama-3.2-3b-instruct:free", expected: "Meta Llama: Llama 3.2 3B Instruct (free)"},
		{input: "meta-llama/llama-3.2-90b-vision-instruct", expected: "Meta Llama: Llama 3.2 90B Vision Instruct"},
		{input: "meta-llama/llama-3.3-70b-instruct", expected: "Meta Llama: Llama 3.3 70B Instruct"},
		{input: "meta-llama/llama-3.3-70b-instruct:free", expected: "Meta Llama: Llama 3.3 70B Instruct (free)"},
		{input: "meta-llama/llama-4-maverick", expected: "Meta Llama: Llama 4 Maverick"},
		{input: "meta-llama/llama-4-scout", expected: "Meta Llama: Llama 4 Scout"},
		{input: "meta-llama/llama-guard-2-8b", expected: "Meta Llama: Llama Guard 2 8B"},
		{input: "meta-llama/llama-guard-3-8b", expected: "Meta Llama: Llama Guard 3 8B"},
		{input: "meta-llama/llama-guard-4-12b", expected: "Meta Llama: Llama Guard 4 12B"},
		{input: "microsoft/phi-3-medium-128k-instruct", expected: "Microsoft: Phi 3 Medium 128K Instruct"},
		{input: "microsoft/phi-3-mini-128k-instruct", expected: "Microsoft: Phi 3 Mini 128K Instruct"},
		{input: "microsoft/phi-3.5-mini-128k-instruct", expected: "Microsoft: Phi 3.5 Mini 128K Instruct"},
		{input: "microsoft/phi-4", expected: "Microsoft: Phi 4"},
		{input: "microsoft/phi-4-multimodal-instruct", expected: "Microsoft: Phi 4 Multimodal Instruct"},
		{input: "microsoft/phi-4-reasoning-plus", expected: "Microsoft: Phi 4 Reasoning Plus"},
		{input: "microsoft/wizardlm-2-8x22b", expected: "Microsoft: WizardLM 2 8x22B"},
		{input: "minimax/minimax-01", expected: "MiniMax: Minimax 01"},
		{input: "minimax/minimax-m1", expected: "MiniMax: Minimax M1"},
		{input: "minimax/minimax-m2", expected: "MiniMax: Minimax M2"},
		{input: "mistralai/codestral-2508", expected: "Mistralai: Codestral 2508"},
		{input: "mistralai/devstral-2512", expected: "Mistralai: Devstral 2512"},
		{input: "mistralai/devstral-2512:free", expected: "Mistralai: Devstral 2512 (free)"},
		{input: "mistralai/devstral-medium", expected: "Mistralai: Devstral Medium"},
		{input: "mistralai/devstral-small", expected: "Mistralai: Devstral Small"},
		{input: "mistralai/devstral-small-2505", expected: "Mistralai: Devstral Small 2505"},
		{input: "mistralai/ministral-14b-2512", expected: "Mistralai: Ministral 14B 2512"},
		{input: "mistralai/ministral-3b", expected: "Mistralai: Ministral 3B"},
		{input: "mistralai/ministral-3b-2512", expected: "Mistralai: Ministral 3B 2512"},
		{input: "mistralai/ministral-8b", expected: "Mistralai: Ministral 8B"},
		{input: "mistralai/ministral-8b-2512", expected: "Mistralai: Ministral 8B 2512"},
		{input: "mistralai/mistral-7b-instruct", expected: "Mistralai: Mistral 7B Instruct"},
		{input: "mistralai/mistral-7b-instruct-v0.1", expected: "Mistralai: Mistral 7B Instruct v0.1"},
		{input: "mistralai/mistral-7b-instruct-v0.2", expected: "Mistralai: Mistral 7B Instruct v0.2"},
		{input: "mistralai/mistral-7b-instruct-v0.3", expected: "Mistralai: Mistral 7B Instruct v0.3"},
		{input: "mistralai/mistral-7b-instruct:free", expected: "Mistralai: Mistral 7B Instruct (free)"},
		{input: "mistralai/mistral-large", expected: "Mistralai: Mistral Large"},
		{input: "mistralai/mistral-large-2407", expected: "Mistralai: Mistral Large 2407"},
		{input: "mistralai/mistral-large-2411", expected: "Mistralai: Mistral Large 2411"},
		{input: "mistralai/mistral-large-2512", expected: "Mistralai: Mistral Large 2512"},
		{input: "mistralai/mistral-medium-3", expected: "Mistralai: Mistral Medium 3"},
		{input: "mistralai/mistral-medium-3.1", expected: "Mistralai: Mistral Medium 3.1"},
		{input: "mistralai/mistral-nemo", expected: "Mistralai: Mistral Nemo"},
		{input: "mistralai/mistral-saba", expected: "Mistralai: Mistral Saba"},
		{input: "mistralai/mistral-small-24b-instruct-2501", expected: "Mistralai: Mistral Small 24B Instruct 2501"},
		{input: "mistralai/mistral-small-3.1-24b-instruct", expected: "Mistralai: Mistral Small 3.1 24B Instruct"},
		{input: "mistralai/mistral-small-3.1-24b-instruct:free", expected: "Mistralai: Mistral Small 3.1 24B Instruct (free)"},
		{input: "mistralai/mistral-small-3.2-24b-instruct", expected: "Mistralai: Mistral Small 3.2 24B Instruct"},
		{input: "mistralai/mistral-small-creative", expected: "Mistralai: Mistral Small Creative"},
		{input: "mistralai/mistral-tiny", expected: "Mistralai: Mistral Tiny"},
		{input: "mistralai/mixtral-8x22b-instruct", expected: "Mistralai: Mixtral 8x22B Instruct"},
		{input: "mistralai/mixtral-8x7b-instruct", expected: "Mistralai: Mixtral 8x7B Instruct"},
		{input: "mistralai/pixtral-12b", expected: "Mistralai: Pixtral 12B"},
		{input: "mistralai/pixtral-large-2411", expected: "Mistralai: Pixtral Large 2411"},
		{input: "mistralai/voxtral-small-24b-2507", expected: "Mistralai: Voxtral Small 24B 2507"},
		{input: "moonshotai/kimi-dev-72b", expected: "Moonshotai: Kimi Dev 72B"},
		{input: "moonshotai/kimi-k2", expected: "Moonshotai: Kimi K2"},
		{input: "moonshotai/kimi-k2-0905", expected: "Moonshotai: Kimi K2 0905"},
		{input: "moonshotai/kimi-k2-0905:exacto", expected: "Moonshotai: Kimi K2 0905 (exacto)"},
		{input: "moonshotai/kimi-k2-thinking", expected: "Moonshotai: Kimi K2 Thinking"},
		{input: "moonshotai/kimi-k2:free", expected: "Moonshotai: Kimi K2 (free)"},
		{input: "morph/morph-v3-fast", expected: "Morph: Morph v3 Fast"},
		{input: "morph/morph-v3-large", expected: "Morph: Morph v3 Large"},
		{input: "neversleep/llama-3.1-lumimaid-8b", expected: "Neversleep: Llama 3.1 Lumimaid 8B"},
		{input: "neversleep/noromaid-20b", expected: "Neversleep: Noromaid 20B"},
		{input: "nex-agi/deepseek-v3.1-nex-n1:free", expected: "Nex Agi: Deepseek v3.1 Nex N1 (free)"},
		{input: "nousresearch/deephermes-3-mistral-24b-preview", expected: "NousResearch: Deephermes 3 Mistral 24B Preview"},
		{input: "nousresearch/hermes-2-pro-llama-3-8b", expected: "NousResearch: Hermes 2 Pro Llama 3 8B"},
		{input: "nousresearch/hermes-3-llama-3.1-405b", expected: "NousResearch: Hermes 3 Llama 3.1 405B"},
		{input: "nousresearch/hermes-3-llama-3.1-405b:free", expected: "NousResearch: Hermes 3 Llama 3.1 405B (free)"},
		{input: "nousresearch/hermes-3-llama-3.1-70b", expected: "NousResearch: Hermes 3 Llama 3.1 70B"},
		{input: "nousresearch/hermes-4-405b", expected: "NousResearch: Hermes 4 405B"},
		{input: "nousresearch/hermes-4-70b", expected: "NousResearch: Hermes 4 70B"},
		{input: "nvidia/llama-3.1-nemotron-70b-instruct", expected: "Nvidia: Llama 3.1 Nemotron 70B Instruct"},
		{input: "nvidia/llama-3.1-nemotron-ultra-253b-v1", expected: "Nvidia: Llama 3.1 Nemotron Ultra 253B v1"},
		{input: "nvidia/llama-3.3-nemotron-super-49b-v1.5", expected: "Nvidia: Llama 3.3 Nemotron Super 49B v1.5"},
		{input: "nvidia/nemotron-3-nano-30b-a3b", expected: "Nvidia: Nemotron 3 Nano 30B A3B"},
		{input: "nvidia/nemotron-3-nano-30b-a3b:free", expected: "Nvidia: Nemotron 3 Nano 30B A3B (free)"},
		{input: "nvidia/nemotron-nano-12b-v2-vl", expected: "Nvidia: Nemotron Nano 12B v2 Vl"},
		{input: "nvidia/nemotron-nano-12b-v2-vl:free", expected: "Nvidia: Nemotron Nano 12B v2 Vl (free)"},
		{input: "nvidia/nemotron-nano-9b-v2", expected: "Nvidia: Nemotron Nano 9B v2"},
		{input: "nvidia/nemotron-nano-9b-v2:free", expected: "Nvidia: Nemotron Nano 9B v2 (free)"},
		{input: "openai/chatgpt-4o-latest", expected: "OpenAI: ChatGPT 4o Latest"},
		{input: "openai/codex-mini", expected: "OpenAI: Codex Mini"},
		{input: "openai/gpt-3.5-turbo", expected: "OpenAI: GPT 3.5 Turbo"},
		{input: "openai/gpt-3.5-turbo-0613", expected: "OpenAI: GPT 3.5 Turbo 0613"},
		{input: "openai/gpt-3.5-turbo-16k", expected: "OpenAI: GPT 3.5 Turbo 16K"},
		{input: "openai/gpt-3.5-turbo-instruct", expected: "OpenAI: GPT 3.5 Turbo Instruct"},
		{input: "openai/gpt-4", expected: "OpenAI: GPT 4"},
		{input: "openai/gpt-4-0314", expected: "OpenAI: GPT 4-0314"},
		{input: "openai/gpt-4-1106-preview", expected: "OpenAI: GPT 4-1106 Preview"},
		{input: "openai/gpt-4-turbo", expected: "OpenAI: GPT 4 Turbo"},
		{input: "openai/gpt-4-turbo-preview", expected: "OpenAI: GPT 4 Turbo Preview"},
		{input: "openai/gpt-4.1", expected: "OpenAI: GPT 4.1"},
		{input: "openai/gpt-4.1-mini", expected: "OpenAI: GPT 4.1 Mini"},
		{input: "openai/gpt-4.1-nano", expected: "OpenAI: GPT 4.1 Nano"},
		{input: "openai/gpt-4o", expected: "OpenAI: GPT 4o"},
		{input: "openai/gpt-4o-2024-05-13", expected: "OpenAI: GPT 4o (2024-05-13)"},
		{input: "openai/gpt-4o-2024-08-06", expected: "OpenAI: GPT 4o (2024-08-06)"},
		{input: "openai/gpt-4o-2024-11-20", expected: "OpenAI: GPT 4o (2024-11-20)"},
		{input: "openai/gpt-4o-audio-preview", expected: "OpenAI: GPT 4o Audio Preview"},
		{input: "openai/gpt-4o-mini", expected: "OpenAI: GPT 4o Mini"},
		{input: "openai/gpt-4o-mini-2024-07-18", expected: "OpenAI: GPT 4o Mini (2024-07-18)"},
		{input: "openai/gpt-4o-mini-search-preview", expected: "OpenAI: GPT 4o Mini Search Preview"},
		{input: "openai/gpt-4o-search-preview", expected: "OpenAI: GPT 4o Search Preview"},
		{input: "openai/gpt-4o:extended", expected: "OpenAI: GPT 4o (extended)"},
		{input: "openai/gpt-5", expected: "OpenAI: GPT 5"},
		{input: "openai/gpt-5-chat", expected: "OpenAI: GPT 5 Chat"},
		{input: "openai/gpt-5-codex", expected: "OpenAI: GPT 5 Codex"},
		{input: "openai/gpt-5-image", expected: "OpenAI: GPT 5 Image"},
		{input: "openai/gpt-5-image-mini", expected: "OpenAI: GPT 5 Image Mini"},
		{input: "openai/gpt-5-mini", expected: "OpenAI: GPT 5 Mini"},
		{input: "openai/gpt-5-nano", expected: "OpenAI: GPT 5 Nano"},
		{input: "openai/gpt-5-pro", expected: "OpenAI: GPT 5 Pro"},
		{input: "openai/gpt-5.1", expected: "OpenAI: GPT 5.1"},
		{input: "openai/gpt-5.1-chat", expected: "OpenAI: GPT 5.1 Chat"},
		{input: "openai/gpt-5.1-codex", expected: "OpenAI: GPT 5.1 Codex"},
		{input: "openai/gpt-5.1-codex-max", expected: "OpenAI: GPT 5.1 Codex Max"},
		{input: "openai/gpt-5.1-codex-mini", expected: "OpenAI: GPT 5.1 Codex Mini"},
		{input: "openai/gpt-5.2", expected: "OpenAI: GPT 5.2"},
		{input: "openai/gpt-5.2-chat", expected: "OpenAI: GPT 5.2 Chat"},
		{input: "openai/gpt-5.2-pro", expected: "OpenAI: GPT 5.2 Pro"},
		{input: "openai/gpt-oss-120b", expected: "OpenAI: GPT OSS 120B"},
		{input: "openai/gpt-oss-120b:exacto", expected: "OpenAI: GPT OSS 120B (exacto)"},
		{input: "openai/gpt-oss-120b:free", expected: "OpenAI: GPT OSS 120B (free)"},
		{input: "openai/gpt-oss-20b", expected: "OpenAI: GPT OSS 20B"},
		{input: "openai/gpt-oss-20b:free", expected: "OpenAI: GPT OSS 20B (free)"},
		{input: "openai/gpt-oss-safeguard-20b", expected: "OpenAI: GPT OSS Safeguard 20B"},
		{input: "openai/o1", expected: "OpenAI: o1"},
		{input: "openai/o1-pro", expected: "OpenAI: o1 Pro"},
		{input: "openai/o3", expected: "OpenAI: o3"},
		{input: "openai/o3-deep-research", expected: "OpenAI: o3 Deep Research"},
		{input: "openai/o3-mini", expected: "OpenAI: o3 Mini"},
		{input: "openai/o3-mini-high", expected: "OpenAI: o3 Mini High"},
		{input: "openai/o3-pro", expected: "OpenAI: o3 Pro"},
		{input: "openai/o4-mini", expected: "OpenAI: o4 Mini"},
		{input: "openai/o4-mini-deep-research", expected: "OpenAI: o4 Mini Deep Research"},
		{input: "openai/o4-mini-high", expected: "OpenAI: o4 Mini High"},
		{input: "opengvlab/internvl3-78b", expected: "Opengvlab: InternVL3 78B"},
		{input: "openrouter/auto", expected: "Openrouter: Auto"},
		{input: "openrouter/bodybuilder", expected: "Openrouter: Bodybuilder"},
		{input: "perplexity/sonar", expected: "Perplexity: Sonar"},
		{input: "perplexity/sonar-deep-research", expected: "Perplexity: Sonar Deep Research"},
		{input: "perplexity/sonar-pro", expected: "Perplexity: Sonar Pro"},
		{input: "perplexity/sonar-pro-search", expected: "Perplexity: Sonar Pro Search"},
		{input: "perplexity/sonar-reasoning", expected: "Perplexity: Sonar Reasoning"},
		{input: "perplexity/sonar-reasoning-pro", expected: "Perplexity: Sonar Reasoning Pro"},
		{input: "prime-intellect/intellect-3", expected: "Prime Intellect: Intellect 3"},
		{input: "qwen/qwen-2.5-72b-instruct", expected: "Qwen: Qwen 2.5 72B Instruct"},
		{input: "qwen/qwen-2.5-7b-instruct", expected: "Qwen: Qwen 2.5 7B Instruct"},
		{input: "qwen/qwen-2.5-coder-32b-instruct", expected: "Qwen: Qwen 2.5 Coder 32B Instruct"},
		{input: "qwen/qwen-2.5-vl-7b-instruct", expected: "Qwen: Qwen 2.5 Vl 7B Instruct"},
		{input: "qwen/qwen-2.5-vl-7b-instruct:free", expected: "Qwen: Qwen 2.5 Vl 7B Instruct (free)"},
		{input: "qwen/qwen-max", expected: "Qwen: Qwen Max"},
		{input: "qwen/qwen-plus", expected: "Qwen: Qwen Plus"},
		{input: "qwen/qwen-plus-2025-07-28", expected: "Qwen: Qwen Plus (2025-07-28)"},
		{input: "qwen/qwen-plus-2025-07-28:thinking", expected: "Qwen: Qwen Plus (2025-07-28) Thinking"},
		{input: "qwen/qwen-turbo", expected: "Qwen: Qwen Turbo"},
		{input: "qwen/qwen-vl-max", expected: "Qwen: Qwen Vl Max"},
		{input: "qwen/qwen-vl-plus", expected: "Qwen: Qwen Vl Plus"},
		{input: "qwen/qwen2.5-coder-7b-instruct", expected: "Qwen: Qwen2.5 Coder 7B Instruct"},
		{input: "qwen/qwen2.5-vl-32b-instruct", expected: "Qwen: Qwen2.5 Vl 32B Instruct"},
		{input: "qwen/qwen2.5-vl-72b-instruct", expected: "Qwen: Qwen2.5 Vl 72B Instruct"},
		{input: "qwen/qwen3-14b", expected: "Qwen: Qwen3 14B"},
		{input: "qwen/qwen3-235b-a22b", expected: "Qwen: Qwen3 235B A22B"},
		{input: "qwen/qwen3-235b-a22b-2507", expected: "Qwen: Qwen3 235B A22B 2507"},
		{input: "qwen/qwen3-235b-a22b-thinking-2507", expected: "Qwen: Qwen3 235B A22B Thinking 2507"},
		{input: "qwen/qwen3-30b-a3b", expected: "Qwen: Qwen3 30B A3B"},
		{input: "qwen/qwen3-30b-a3b-instruct-2507", expected: "Qwen: Qwen3 30B A3B Instruct 2507"},
		{input: "qwen/qwen3-30b-a3b-thinking-2507", expected: "Qwen: Qwen3 30B A3B Thinking 2507"},
		{input: "qwen/qwen3-32b", expected: "Qwen: Qwen3 32B"},
		{input: "qwen/qwen3-4b:free", expected: "Qwen: Qwen3 4B (free)"},
		{input: "qwen/qwen3-8b", expected: "Qwen: Qwen3 8B"},
		{input: "qwen/qwen3-coder", expected: "Qwen: Qwen3 Coder"},
		{input: "qwen/qwen3-coder-30b-a3b-instruct", expected: "Qwen: Qwen3 Coder 30B A3B Instruct"},
		{input: "qwen/qwen3-coder-flash", expected: "Qwen: Qwen3 Coder Flash"},
		{input: "qwen/qwen3-coder-plus", expected: "Qwen: Qwen3 Coder Plus"},
		{input: "qwen/qwen3-coder:exacto", expected: "Qwen: Qwen3 Coder (exacto)"},
		{input: "qwen/qwen3-coder:free", expected: "Qwen: Qwen3 Coder (free)"},
		{input: "qwen/qwen3-max", expected: "Qwen: Qwen3 Max"},
		{input: "qwen/qwen3-next-80b-a3b-instruct", expected: "Qwen: Qwen3 Next 80B A3B Instruct"},
		{input: "qwen/qwen3-next-80b-a3b-thinking", expected: "Qwen: Qwen3 Next 80B A3B Thinking"},
		{input: "qwen/qwen3-vl-235b-a22b-instruct", expected: "Qwen: Qwen3 Vl 235B A22B Instruct"},
		{input: "qwen/qwen3-vl-235b-a22b-thinking", expected: "Qwen: Qwen3 Vl 235B A22B Thinking"},
		{input: "qwen/qwen3-vl-30b-a3b-instruct", expected: "Qwen: Qwen3 Vl 30B A3B Instruct"},
		{input: "qwen/qwen3-vl-30b-a3b-thinking", expected: "Qwen: Qwen3 Vl 30B A3B Thinking"},
		{input: "qwen/qwen3-vl-32b-instruct", expected: "Qwen: Qwen3 Vl 32B Instruct"},
		{input: "qwen/qwen3-vl-8b-instruct", expected: "Qwen: Qwen3 Vl 8B Instruct"},
		{input: "qwen/qwen3-vl-8b-thinking", expected: "Qwen: Qwen3 Vl 8B Thinking"},
		{input: "qwen/qwq-32b", expected: "Qwen: Qwq 32B"},
		{input: "raifle/sorcererlm-8x22b", expected: "Raifle: SorcererLM 8x22B"},
		{input: "relace/relace-apply-3", expected: "Relace: Relace Apply 3"},
		{input: "relace/relace-search", expected: "Relace: Relace Search"},
		{input: "sao10k/l3-euryale-70b", expected: "Sao10k: L3 Euryale 70B"},
		{input: "sao10k/l3-lunaris-8b", expected: "Sao10k: L3 Lunaris 8B"},
		{input: "sao10k/l3.1-70b-hanami-x1", expected: "Sao10k: L3.1 70B Hanami X1"},
		{input: "sao10k/l3.1-euryale-70b", expected: "Sao10k: L3.1 Euryale 70B"},
		{input: "sao10k/l3.3-euryale-70b", expected: "Sao10k: L3.3 Euryale 70B"},
		{input: "stepfun-ai/step3", expected: "StepFun: Step3"},
		{input: "switchpoint/router", expected: "Switchpoint: Router"},
		{input: "tencent/hunyuan-a13b-instruct", expected: "Tencent: Hunyuan A13B Instruct"},
		{input: "thedrummer/cydonia-24b-v4.1", expected: "Thedrummer: Cydonia 24B v4.1"},
		{input: "thedrummer/rocinante-12b", expected: "Thedrummer: Rocinante 12B"},
		{input: "thedrummer/skyfall-36b-v2", expected: "Thedrummer: Skyfall 36B v2"},
		{input: "thedrummer/unslopnemo-12b", expected: "Thedrummer: Unslopnemo 12B"},
		{input: "thudm/glm-4.1v-9b-thinking", expected: "Thudm: GLM 4.1V 9B Thinking"},
		{input: "tngtech/deepseek-r1t-chimera", expected: "Tngtech: Deepseek R1T Chimera"},
		{input: "tngtech/deepseek-r1t-chimera:free", expected: "Tngtech: Deepseek R1T Chimera (free)"},
		{input: "tngtech/deepseek-r1t2-chimera", expected: "Tngtech: Deepseek R1T2 Chimera"},
		{input: "tngtech/deepseek-r1t2-chimera:free", expected: "Tngtech: Deepseek R1T2 Chimera (free)"},
		{input: "tngtech/tng-r1t-chimera", expected: "Tngtech: Tng R1T Chimera"},
		{input: "tngtech/tng-r1t-chimera:free", expected: "Tngtech: Tng R1T Chimera (free)"},
		{input: "undi95/remm-slerp-l2-13b", expected: "Undi95: Remm Slerp L2 13B"},
		{input: "x-ai/grok-3", expected: "xAI: Grok 3"},
		{input: "x-ai/grok-3-beta", expected: "xAI: Grok 3 Beta"},
		{input: "x-ai/grok-3-mini", expected: "xAI: Grok 3 Mini"},
		{input: "x-ai/grok-3-mini-beta", expected: "xAI: Grok 3 Mini Beta"},
		{input: "x-ai/grok-4", expected: "xAI: Grok 4"},
		{input: "x-ai/grok-4-fast", expected: "xAI: Grok 4 Fast"},
		{input: "x-ai/grok-4.1-fast", expected: "xAI: Grok 4.1 Fast"},
		{input: "x-ai/grok-code-fast-1", expected: "xAI: Grok Code Fast 1"},
		{input: "xiaomi/mimo-v2-flash:free", expected: "Xiaomi: Mimo v2 Flash (free)"},
		{input: "z-ai/glm-4-32b", expected: "Z.AI: GLM 4 32B"},
		{input: "z-ai/glm-4.5", expected: "Z.AI: GLM 4.5"},
		{input: "z-ai/glm-4.5-air", expected: "Z.AI: GLM 4.5 Air"},
		{input: "z-ai/glm-4.5-air:free", expected: "Z.AI: GLM 4.5 Air (free)"},
		{input: "z-ai/glm-4.5v", expected: "Z.AI: GLM 4.5V"},
		{input: "z-ai/glm-4.6", expected: "Z.AI: GLM 4.6"},
		{input: "z-ai/glm-4.6:exacto", expected: "Z.AI: GLM 4.6 (exacto)"},
		{input: "z-ai/glm-4.6v", expected: "Z.AI: GLM 4.6V"},
		{input: "O1-2025-12-11", expected: "o1 (2025-12-11)"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatModelTitle(tt.input))
		})
	}
}

func Test_ModelTitle_RealTitleData(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "AI21: Jamba Large 1.7", expected: "AI21: Jamba Large 1.7"},
		{input: "AI21: Jamba Mini 1.7", expected: "AI21: Jamba Mini 1.7"},
		{input: "AionLabs: Aion-1.0", expected: "AionLabs: Aion 1.0"},
		{input: "AionLabs: Aion-1.0-Mini", expected: "AionLabs: Aion 1.0 Mini"},
		{input: "AionLabs: Aion-RP 1.0 (8B)", expected: "AionLabs: Aion RP 1.0 (8B)"},
		{input: "AlfredPros: CodeLLaMa 7B Instruct Solidity", expected: "AlfredPros: CodeLLaMa 7B Instruct Solidity"},
		{input: "AllenAI: Olmo 2 32B Instruct", expected: "AllenAI: Olmo 2 32B Instruct"},
		{input: "AllenAI: Olmo 3 32B Think (free)", expected: "AllenAI: Olmo 3 32B Think (free)"},
		{input: "AllenAI: Olmo 3 7B Instruct", expected: "AllenAI: Olmo 3 7B Instruct"},
		{input: "AllenAI: Olmo 3 7B Think", expected: "AllenAI: Olmo 3 7B Think"},
		{input: "AllenAI: Olmo 3.1 32B Think (free)", expected: "AllenAI: Olmo 3.1 32B Think (free)"},
		{input: "Amazon: Nova 2 Lite", expected: "Amazon: Nova 2 Lite"},
		{input: "Amazon: Nova Lite 1.0", expected: "Amazon: Nova Lite 1.0"},
		{input: "Amazon: Nova Micro 1.0", expected: "Amazon: Nova Micro 1.0"},
		{input: "Amazon: Nova Premier 1.0", expected: "Amazon: Nova Premier 1.0"},
		{input: "Amazon: Nova Pro 1.0", expected: "Amazon: Nova Pro 1.0"},
		{input: "Anthropic: Claude 3 Haiku", expected: "Anthropic: Claude 3 Haiku"},
		{input: "Anthropic: Claude 3 Opus", expected: "Anthropic: Claude 3 Opus"},
		{input: "Anthropic: Claude 3.5 Haiku", expected: "Anthropic: Claude 3.5 Haiku"},
		{input: "Anthropic: Claude 3.5 Haiku (2024-10-22)", expected: "Anthropic: Claude 3.5 Haiku (2024-10-22)"},
		{input: "Anthropic: Claude 3.5 Sonnet", expected: "Anthropic: Claude 3.5 Sonnet"},
		{input: "Anthropic: Claude 3.7 Sonnet", expected: "Anthropic: Claude 3.7 Sonnet"},
		{input: "Anthropic: Claude 3.7 Sonnet (thinking)", expected: "Anthropic: Claude 3.7 Sonnet (thinking)"},
		{input: "Anthropic: Claude Haiku 4.5", expected: "Anthropic: Claude Haiku 4.5"},
		{input: "Anthropic: Claude Opus 4", expected: "Anthropic: Claude Opus 4"},
		{input: "Anthropic: Claude Opus 4.1", expected: "Anthropic: Claude Opus 4.1"},
		{input: "Anthropic: Claude Opus 4.5", expected: "Anthropic: Claude Opus 4.5"},
		{input: "Anthropic: Claude Sonnet 4", expected: "Anthropic: Claude Sonnet 4"},
		{input: "Anthropic: Claude Sonnet 4.5", expected: "Anthropic: Claude Sonnet 4.5"},
		{input: "ArceeAI: Coder Large", expected: "ArceeAI: Coder Large"},
		{input: "ArceeAI: Maestro Reasoning", expected: "ArceeAI: Maestro Reasoning"},
		{input: "ArceeAI: Spotlight", expected: "ArceeAI: Spotlight"},
		{input: "ArceeAI: Trinity Mini", expected: "ArceeAI: Trinity Mini"},
		{input: "ArceeAI: Trinity Mini (free)", expected: "ArceeAI: Trinity Mini (free)"},
		{input: "ArceeAI: Virtuoso Large", expected: "ArceeAI: Virtuoso Large"},
		{input: "ArliAI: QwQ 32B RpR v1", expected: "ArliAI: QwQ 32B RpR v1"},
		{input: "Auto Router", expected: "Auto Router"},
		{input: "Baidu: ERNIE 4.5 21B A3B", expected: "Baidu: ERNIE 4.5 21B A3B"},
		{input: "Baidu: ERNIE 4.5 21B A3B Thinking", expected: "Baidu: ERNIE 4.5 21B A3B Thinking"},
		{input: "Baidu: ERNIE 4.5 300B A47B", expected: "Baidu: ERNIE 4.5 300B A47B"},
		{input: "Baidu: ERNIE 4.5 VL 28B A3B", expected: "Baidu: ERNIE 4.5 VL 28B A3B"},
		{input: "Baidu: ERNIE 4.5 VL 424B A47B", expected: "Baidu: ERNIE 4.5 VL 424B A47B"},
		{input: "ByteDance: UI-TARS 7B", expected: "ByteDance: UI TARS 7B"},
		{input: "Body Builder (beta)", expected: "Body Builder (beta)"},
		{input: "Cogito V2 Preview Llama 109B", expected: "Cogito v2 Preview Llama 109B"},
		{input: "Cohere: Command A", expected: "Cohere: Command A"},
		{input: "Cohere: Command R (08-2024)", expected: "Cohere: Command R (2024-08)"},
		{input: "Cohere: Command R+ (08-2024)", expected: "Cohere: Command R+ (2024-08)"},
		{input: "Cohere: Command R7B (12-2024)", expected: "Cohere: Command R7B (2024-12)"},
		{input: "Deep Cogito: Cogito V2 Preview Llama 405B", expected: "Deep Cogito: Cogito v2 Preview Llama 405B"},
		{input: "Deep Cogito: Cogito V2 Preview Llama 70B", expected: "Deep Cogito: Cogito v2 Preview Llama 70B"},
		{input: "Deep Cogito: Cogito v2.1 671B", expected: "Deep Cogito: Cogito v2.1 671B"},
		{input: "DeepSeek: DeepSeek Prover V2", expected: "DeepSeek: DeepSeek Prover v2"},
		{input: "DeepSeek: DeepSeek R1 0528 Qwen3 8B", expected: "DeepSeek: DeepSeek R1 0528 Qwen3 8B"},
		{input: "DeepSeek: DeepSeek V3", expected: "DeepSeek: DeepSeek v3"},
		{input: "DeepSeek: DeepSeek V3 0324", expected: "DeepSeek: DeepSeek v3 0324"},
		{input: "DeepSeek: DeepSeek V3.1", expected: "DeepSeek: DeepSeek v3.1"},
		{input: "DeepSeek: DeepSeek V3.1 Terminus", expected: "DeepSeek: DeepSeek v3.1 Terminus"},
		{input: "DeepSeek: DeepSeek V3.1 Terminus (exacto)", expected: "DeepSeek: DeepSeek v3.1 Terminus (exacto)"},
		{input: "DeepSeek: DeepSeek V3.2", expected: "DeepSeek: DeepSeek v3.2"},
		{input: "DeepSeek: DeepSeek V3.2 Exp", expected: "DeepSeek: DeepSeek v3.2 Exp"},
		{input: "DeepSeek: DeepSeek V3.2 Speciale", expected: "DeepSeek: DeepSeek v3.2 Speciale"},
		{input: "DeepSeek: R1", expected: "DeepSeek: R1"},
		{input: "DeepSeek: R1 0528", expected: "DeepSeek: R1 0528"},
		{input: "DeepSeek: R1 0528 (free)", expected: "DeepSeek: R1 0528 (free)"},
		{input: "DeepSeek: R1 Distill Llama 70B", expected: "DeepSeek: R1 Distill Llama 70B"},
		{input: "DeepSeek: R1 Distill Qwen 14B", expected: "DeepSeek: R1 Distill Qwen 14B"},
		{input: "DeepSeek: R1 Distill Qwen 32B", expected: "DeepSeek: R1 Distill Qwen 32B"},
		{input: "EleutherAI: Llemma 7b", expected: "EleutherAI: Llemma 7B"},
		{input: "EssentialAI: Rnj 1 Instruct", expected: "EssentialAI: Rnj 1 Instruct"},
		{input: "Goliath 120B", expected: "Goliath 120B"},
		{input: "Google: Gemini 2.0 Flash", expected: "Google: Gemini 2.0 Flash"},
		{input: "Google: Gemini 2.0 Flash Experimental (free)", expected: "Google: Gemini 2.0 Flash Experimental (free)"},
		{input: "Google: Gemini 2.0 Flash Lite", expected: "Google: Gemini 2.0 Flash Lite"},
		{input: "Google: Gemini 2.5 Flash", expected: "Google: Gemini 2.5 Flash"},
		{input: "Google: Gemini 2.5 Flash Image (Nano Banana)", expected: "Google: Gemini 2.5 Flash Image (Nano Banana)"},
		{input: "Google: Gemini 2.5 Flash Image Preview (Nano Banana)", expected: "Google: Gemini 2.5 Flash Image Preview (Nano Banana)"},
		{input: "Google: Gemini 2.5 Flash Lite", expected: "Google: Gemini 2.5 Flash Lite"},
		{input: "Google: Gemini 2.5 Flash Lite Preview 09-2025", expected: "Google: Gemini 2.5 Flash Lite Preview (2025-09)"},
		{input: "Google: Gemini 2.5 Flash Preview 09-2025", expected: "Google: Gemini 2.5 Flash Preview (2025-09)"},
		{input: "Google: Gemini 2.5 Pro", expected: "Google: Gemini 2.5 Pro"},
		{input: "Google: Gemini 2.5 Pro Preview 05-06", expected: "Google: Gemini 2.5 Pro Preview 05-06"},
		{input: "Google: Gemini 2.5 Pro Preview 06-05", expected: "Google: Gemini 2.5 Pro Preview 06-05"},
		{input: "Google: Gemini 3 Flash Preview", expected: "Google: Gemini 3 Flash Preview"},
		{input: "Google: Gemini 3 Pro Preview", expected: "Google: Gemini 3 Pro Preview"},
		{input: "Google: Gemma 2 27B", expected: "Google: Gemma 2 27B"},
		{input: "Google: Gemma 2 9B", expected: "Google: Gemma 2 9B"},
		{input: "Google: Gemma 3 12B", expected: "Google: Gemma 3 12B"},
		{input: "Google: Gemma 3 12B (free)", expected: "Google: Gemma 3 12B (free)"},
		{input: "Google: Gemma 3 27B", expected: "Google: Gemma 3 27B"},
		{input: "Google: Gemma 3 27B (free)", expected: "Google: Gemma 3 27B (free)"},
		{input: "Google: Gemma 3 4B", expected: "Google: Gemma 3 4B"},
		{input: "Google: Gemma 3 4B (free)", expected: "Google: Gemma 3 4B (free)"},
		{input: "Google: Gemma 3n 2B (free)", expected: "Google: Gemma 3n 2B (free)"},
		{input: "Google: Gemma 3n 4B", expected: "Google: Gemma 3n 4B"},
		{input: "Google: Gemma 3n 4B (free)", expected: "Google: Gemma 3n 4B (free)"},
		{input: "Google: Nano Banana Pro (Gemini 3 Pro Image Preview)", expected: "Google: Nano Banana Pro (Gemini 3 Pro Image Preview)"},
		{input: "IBM: Granite 4.0 Micro", expected: "IBM: Granite 4.0 Micro"},
		{input: "Inception: Mercury", expected: "Inception: Mercury"},
		{input: "Inception: Mercury Coder", expected: "Inception: Mercury Coder"},
		{input: "Inflection: Inflection 3 Pi", expected: "Inflection: Inflection 3 Pi"},
		{input: "Inflection: Inflection 3 Productivity", expected: "Inflection: Inflection 3 Productivity"},
		{input: "Kwaipilot: KAT-Coder-Pro V1 (free)", expected: "Kwaipilot: KAT Coder Pro v1 (free)"},
		{input: "LiquidAI/LFM2-2.6B", expected: "LiquidAI: LFM2 2.6B"},
		{input: "LiquidAI/LFM2-8B-A1B", expected: "LiquidAI: LFM2 8B A1B"},
		{input: "Llama Guard 3 8B", expected: "Llama Guard 3 8B"},
		{input: "Magnum v4 72B", expected: "Magnum v4 72B"},
		{input: "Mancer: Weaver (alpha)", expected: "Mancer: Weaver (alpha)"},
		{input: "Meituan: LongCat Flash Chat", expected: "Meituan: LongCat Flash Chat"},
		{input: "Meta: Llama 3 70B Instruct", expected: "Meta: Llama 3 70B Instruct"},
		{input: "Meta: Llama 3 8B Instruct", expected: "Meta: Llama 3 8B Instruct"},
		{input: "Meta: Llama 3.1 405B (base)", expected: "Meta: Llama 3.1 405B (base)"},
		{input: "Meta: Llama 3.1 405B Instruct", expected: "Meta: Llama 3.1 405B Instruct"},
		{input: "Meta: Llama 3.1 405B Instruct (free)", expected: "Meta: Llama 3.1 405B Instruct (free)"},
		{input: "Meta: Llama 3.1 70B Instruct", expected: "Meta: Llama 3.1 70B Instruct"},
		{input: "Meta: Llama 3.1 8B Instruct", expected: "Meta: Llama 3.1 8B Instruct"},
		{input: "Meta: Llama 3.2 11B Vision Instruct", expected: "Meta: Llama 3.2 11B Vision Instruct"},
		{input: "Meta: Llama 3.2 1B Instruct", expected: "Meta: Llama 3.2 1B Instruct"},
		{input: "Meta: Llama 3.2 3B Instruct", expected: "Meta: Llama 3.2 3B Instruct"},
		{input: "Meta: Llama 3.2 3B Instruct (free)", expected: "Meta: Llama 3.2 3B Instruct (free)"},
		{input: "Meta: Llama 3.2 90B Vision Instruct", expected: "Meta: Llama 3.2 90B Vision Instruct"},
		{input: "Meta: Llama 3.3 70B Instruct", expected: "Meta: Llama 3.3 70B Instruct"},
		{input: "Meta: Llama 3.3 70B Instruct (free)", expected: "Meta: Llama 3.3 70B Instruct (free)"},
		{input: "Meta: Llama 4 Maverick", expected: "Meta: Llama 4 Maverick"},
		{input: "Meta: Llama 4 Scout", expected: "Meta: Llama 4 Scout"},
		{input: "Meta: Llama Guard 4 12B", expected: "Meta: Llama Guard 4 12B"},
		{input: "Meta: LlamaGuard 2 8B", expected: "Meta: LlamaGuard 2 8B"},
		{input: "Microsoft: Phi 4", expected: "Microsoft: Phi 4"},
		{input: "Microsoft: Phi 4 Multimodal Instruct", expected: "Microsoft: Phi 4 Multimodal Instruct"},
		{input: "Microsoft: Phi 4 Reasoning Plus", expected: "Microsoft: Phi 4 Reasoning Plus"},
		{input: "Microsoft: Phi-3 Medium 128K Instruct", expected: "Microsoft: Phi 3 Medium 128K Instruct"},
		{input: "Microsoft: Phi-3 Mini 128K Instruct", expected: "Microsoft: Phi 3 Mini 128K Instruct"},
		{input: "Microsoft: Phi-3.5 Mini 128K Instruct", expected: "Microsoft: Phi 3.5 Mini 128K Instruct"},
		{input: "MiniMax: MiniMax M1", expected: "MiniMax: MiniMax M1"},
		{input: "MiniMax: MiniMax M2", expected: "MiniMax: MiniMax M2"},
		{input: "MiniMax: MiniMax-01", expected: "MiniMax: MiniMax 01"},
		{input: "Mistral Large", expected: "Mistral Large"},
		{input: "Mistral Large 2407", expected: "Mistral Large 2407"},
		{input: "Mistral Large 2411", expected: "Mistral Large 2411"},
		{input: "Mistral Tiny", expected: "Mistral Tiny"},
		{input: "Mistral: Codestral 2508", expected: "Mistral: Codestral 2508"},
		{input: "Mistral: Devstral 2 2512", expected: "Mistral: Devstral 2 2512"},
		{input: "Mistral: Devstral 2 2512 (free)", expected: "Mistral: Devstral 2 2512 (free)"},
		{input: "Mistral: Devstral Medium", expected: "Mistral: Devstral Medium"},
		{input: "Mistral: Devstral Small 1.1", expected: "Mistral: Devstral Small 1.1"},
		{input: "Mistral: Devstral Small 2505", expected: "Mistral: Devstral Small 2505"},
		{input: "Mistral: Ministral 3 14B 2512", expected: "Mistral: Ministral 3 14B 2512"},
		{input: "Mistral: Ministral 3 3B 2512", expected: "Mistral: Ministral 3 3B 2512"},
		{input: "Mistral: Ministral 3 8B 2512", expected: "Mistral: Ministral 3 8B 2512"},
		{input: "Mistral: Ministral 3B", expected: "Mistral: Ministral 3B"},
		{input: "Mistral: Ministral 8B", expected: "Mistral: Ministral 8B"},
		{input: "Mistral: Mistral 7B Instruct", expected: "Mistral: Mistral 7B Instruct"},
		{input: "Mistral: Mistral 7B Instruct (free)", expected: "Mistral: Mistral 7B Instruct (free)"},
		{input: "Mistral: Mistral 7B Instruct v0.1", expected: "Mistral: Mistral 7B Instruct v0.1"},
		{input: "Mistral: Mistral 7B Instruct v0.2", expected: "Mistral: Mistral 7B Instruct v0.2"},
		{input: "Mistral: Mistral 7B Instruct v0.3", expected: "Mistral: Mistral 7B Instruct v0.3"},
		{input: "Mistral: Mistral Large 3 2512", expected: "Mistral: Mistral Large 3 2512"},
		{input: "Mistral: Mistral Medium 3", expected: "Mistral: Mistral Medium 3"},
		{input: "Mistral: Mistral Medium 3.1", expected: "Mistral: Mistral Medium 3.1"},
		{input: "Mistral: Mistral Nemo", expected: "Mistral: Mistral Nemo"},
		{input: "Mistral: Mistral Small 3", expected: "Mistral: Mistral Small 3"},
		{input: "Mistral: Mistral Small 3.1 24B", expected: "Mistral: Mistral Small 3.1 24B"},
		{input: "Mistral: Mistral Small 3.1 24B (free)", expected: "Mistral: Mistral Small 3.1 24B (free)"},
		{input: "Mistral: Mistral Small 3.2 24B", expected: "Mistral: Mistral Small 3.2 24B"},
		{input: "Mistral: Mistral Small Creative", expected: "Mistral: Mistral Small Creative"},
		{input: "Mistral: Mixtral 8x22B Instruct", expected: "Mistral: Mixtral 8x22B Instruct"},
		{input: "Mistral: Mixtral 8x7B Instruct", expected: "Mistral: Mixtral 8x7B Instruct"},
		{input: "Mistral: Pixtral 12B", expected: "Mistral: Pixtral 12B"},
		{input: "Mistral: Pixtral Large 2411", expected: "Mistral: Pixtral Large 2411"},
		{input: "Mistral: Saba", expected: "Mistral: Saba"},
		{input: "Mistral: Voxtral Small 24B 2507", expected: "Mistral: Voxtral Small 24B 2507"},
		{input: "MythoMax 13B", expected: "MythoMax 13B"},
		{input: "MoonshotAI: Kimi Dev 72B", expected: "MoonshotAI: Kimi Dev 72B"},
		{input: "MoonshotAI: Kimi K2 0711", expected: "MoonshotAI: Kimi K2 0711"},
		{input: "MoonshotAI: Kimi K2 0711 (free)", expected: "MoonshotAI: Kimi K2 0711 (free)"},
		{input: "MoonshotAI: Kimi K2 0905", expected: "MoonshotAI: Kimi K2 0905"},
		{input: "MoonshotAI: Kimi K2 0905 (exacto)", expected: "MoonshotAI: Kimi K2 0905 (exacto)"},
		{input: "MoonshotAI: Kimi K2 Thinking", expected: "MoonshotAI: Kimi K2 Thinking"},
		{input: "Morph: Morph V3 Fast", expected: "Morph: Morph v3 Fast"},
		{input: "Morph: Morph V3 Large", expected: "Morph: Morph v3 Large"},
		{input: "NeverSleep: Lumimaid v0.2 8B", expected: "NeverSleep: Lumimaid v0.2 8B"},
		{input: "Nex AGI: DeepSeek V3.1 Nex N1 (free)", expected: "Nex AGI: DeepSeek v3.1 Nex N1 (free)"},
		{input: "Noromaid 20B", expected: "Noromaid 20B"},
		{input: "Nous: DeepHermes 3 Mistral 24B Preview", expected: "Nous: DeepHermes 3 Mistral 24B Preview"},
		{input: "Nous: Hermes 3 405B Instruct", expected: "Nous: Hermes 3 405B Instruct"},
		{input: "Nous: Hermes 3 405B Instruct (free)", expected: "Nous: Hermes 3 405B Instruct (free)"},
		{input: "Nous: Hermes 3 70B Instruct", expected: "Nous: Hermes 3 70B Instruct"},
		{input: "Nous: Hermes 4 405B", expected: "Nous: Hermes 4 405B"},
		{input: "Nous: Hermes 4 70B", expected: "Nous: Hermes 4 70B"},
		{input: "NousResearch: Hermes 2 Pro - Llama-3 8B", expected: "NousResearch: Hermes 2 Pro - Llama 3 8B"},
		{input: "NVIDIA: Llama 3.1 Nemotron 70B Instruct", expected: "NVIDIA: Llama 3.1 Nemotron 70B Instruct"},
		{input: "NVIDIA: Llama 3.1 Nemotron Ultra 253B v1", expected: "NVIDIA: Llama 3.1 Nemotron Ultra 253B v1"},
		{input: "NVIDIA: Llama 3.3 Nemotron Super 49B V1.5", expected: "NVIDIA: Llama 3.3 Nemotron Super 49B v1.5"},
		{input: "NVIDIA: Nemotron 3 Nano 30B A3B", expected: "NVIDIA: Nemotron 3 Nano 30B A3B"},
		{input: "NVIDIA: Nemotron 3 Nano 30B A3B (free)", expected: "NVIDIA: Nemotron 3 Nano 30B A3B (free)"},
		{input: "NVIDIA: Nemotron Nano 12B 2 VL", expected: "NVIDIA: Nemotron Nano 12B 2 VL"},
		{input: "NVIDIA: Nemotron Nano 12B 2 VL (free)", expected: "NVIDIA: Nemotron Nano 12B 2 VL (free)"},
		{input: "NVIDIA: Nemotron Nano 9B V2", expected: "NVIDIA: Nemotron Nano 9B v2"},
		{input: "NVIDIA: Nemotron Nano 9B V2 (free)", expected: "NVIDIA: Nemotron Nano 9B v2 (free)"},
		{input: "OpenAI: ChatGPT-4o", expected: "OpenAI: ChatGPT 4o"},
		{input: "OpenAI: Codex Mini", expected: "OpenAI: Codex Mini"},
		{input: "OpenAI: GPT-3.5 Turbo", expected: "OpenAI: GPT 3.5 Turbo"},
		{input: "OpenAI: GPT-3.5 Turbo (older v0613)", expected: "OpenAI: GPT 3.5 Turbo (older v0613)"},
		{input: "OpenAI: GPT-3.5 Turbo 16k", expected: "OpenAI: GPT 3.5 Turbo 16K"},
		{input: "OpenAI: GPT-3.5 Turbo Instruct", expected: "OpenAI: GPT 3.5 Turbo Instruct"},
		{input: "OpenAI: GPT-4", expected: "OpenAI: GPT 4"},
		{input: "OpenAI: GPT-4 (older v0314)", expected: "OpenAI: GPT 4 (older v0314)"},
		{input: "OpenAI: GPT-4 Turbo", expected: "OpenAI: GPT 4 Turbo"},
		{input: "OpenAI: GPT-4 Turbo (older v1106)", expected: "OpenAI: GPT 4 Turbo (older v1106)"},
		{input: "OpenAI: GPT-4 Turbo Preview", expected: "OpenAI: GPT 4 Turbo Preview"},
		{input: "OpenAI: GPT-4.1", expected: "OpenAI: GPT 4.1"},
		{input: "OpenAI: GPT-4.1 Mini", expected: "OpenAI: GPT 4.1 Mini"},
		{input: "OpenAI: GPT-4.1 Nano", expected: "OpenAI: GPT 4.1 Nano"},
		{input: "OpenAI: GPT-4o", expected: "OpenAI: GPT 4o"},
		{input: "OpenAI: GPT-4o (2024-05-13)", expected: "OpenAI: GPT 4o (2024-05-13)"},
		{input: "OpenAI: GPT-4o (2024-08-06)", expected: "OpenAI: GPT 4o (2024-08-06)"},
		{input: "OpenAI: GPT-4o (2024-11-20)", expected: "OpenAI: GPT 4o (2024-11-20)"},
		{input: "OpenAI: GPT-4o (extended)", expected: "OpenAI: GPT 4o (extended)"},
		{input: "OpenAI: GPT-4o Audio", expected: "OpenAI: GPT 4o Audio"},
		{input: "OpenAI: GPT-4o Search Preview", expected: "OpenAI: GPT 4o Search Preview"},
		{input: "OpenAI: GPT-4o-mini", expected: "OpenAI: GPT 4o Mini"},
		{input: "OpenAI: GPT-4o-mini (2024-07-18)", expected: "OpenAI: GPT 4o Mini (2024-07-18)"},
		{input: "OpenAI: GPT-4o-mini Search Preview", expected: "OpenAI: GPT 4o Mini Search Preview"},
		{input: "OpenAI: GPT-5", expected: "OpenAI: GPT 5"},
		{input: "OpenAI: GPT-5 Chat", expected: "OpenAI: GPT 5 Chat"},
		{input: "OpenAI: GPT-5 Codex", expected: "OpenAI: GPT 5 Codex"},
		{input: "OpenAI: GPT-5 Image", expected: "OpenAI: GPT 5 Image"},
		{input: "OpenAI: GPT-5 Image Mini", expected: "OpenAI: GPT 5 Image Mini"},
		{input: "OpenAI: GPT-5 Mini", expected: "OpenAI: GPT 5 Mini"},
		{input: "OpenAI: GPT-5 Nano", expected: "OpenAI: GPT 5 Nano"},
		{input: "OpenAI: GPT-5 Pro", expected: "OpenAI: GPT 5 Pro"},
		{input: "OpenAI: GPT-5.1", expected: "OpenAI: GPT 5.1"},
		{input: "OpenAI: GPT-5.1 Chat", expected: "OpenAI: GPT 5.1 Chat"},
		{input: "OpenAI: GPT-5.1-Codex", expected: "OpenAI: GPT 5.1 Codex"},
		{input: "OpenAI: GPT-5.1-Codex-Max", expected: "OpenAI: GPT 5.1 Codex Max"},
		{input: "OpenAI: GPT-5.1-Codex-Mini", expected: "OpenAI: GPT 5.1 Codex Mini"},
		{input: "OpenAI: GPT-5.2", expected: "OpenAI: GPT 5.2"},
		{input: "OpenAI: GPT-5.2 Chat", expected: "OpenAI: GPT 5.2 Chat"},
		{input: "OpenAI: GPT-5.2 Pro", expected: "OpenAI: GPT 5.2 Pro"},
		{input: "OpenAI: gpt-oss-120b", expected: "OpenAI: GPT OSS 120B"},
		{input: "OpenAI: gpt-oss-120b (exacto)", expected: "OpenAI: GPT OSS 120B (exacto)"},
		{input: "OpenAI: gpt-oss-120b (free)", expected: "OpenAI: GPT OSS 120B (free)"},
		{input: "OpenAI: gpt-oss-20b", expected: "OpenAI: GPT OSS 20B"},
		{input: "OpenAI: gpt-oss-20b (free)", expected: "OpenAI: GPT OSS 20B (free)"},
		{input: "OpenAI: gpt-oss-safeguard-20b", expected: "OpenAI: GPT OSS Safeguard 20B"},
		{input: "OpenAI: o1", expected: "OpenAI: o1"},
		{input: "OpenAI: o1-pro", expected: "OpenAI: o1 Pro"},
		{input: "OpenAI: o3", expected: "OpenAI: o3"},
		{input: "OpenAI: o3 Deep Research", expected: "OpenAI: o3 Deep Research"},
		{input: "OpenAI: o3 Mini", expected: "OpenAI: o3 Mini"},
		{input: "OpenAI: o3 Mini High", expected: "OpenAI: o3 Mini High"},
		{input: "OpenAI: o3 Pro", expected: "OpenAI: o3 Pro"},
		{input: "OpenAI: o4 Mini", expected: "OpenAI: o4 Mini"},
		{input: "OpenAI: o4 Mini Deep Research", expected: "OpenAI: o4 Mini Deep Research"},
		{input: "OpenAI: o4 Mini High", expected: "OpenAI: o4 Mini High"},
		{input: "OpenGVLab: InternVL3 78B", expected: "OpenGVLab: InternVL3 78B"},
		{input: "Perplexity: Sonar", expected: "Perplexity: Sonar"},
		{input: "Perplexity: Sonar Deep Research", expected: "Perplexity: Sonar Deep Research"},
		{input: "Perplexity: Sonar Pro", expected: "Perplexity: Sonar Pro"},
		{input: "Perplexity: Sonar Pro Search", expected: "Perplexity: Sonar Pro Search"},
		{input: "Perplexity: Sonar Reasoning", expected: "Perplexity: Sonar Reasoning"},
		{input: "Perplexity: Sonar Reasoning Pro", expected: "Perplexity: Sonar Reasoning Pro"},
		{input: "Prime Intellect: INTELLECT-3", expected: "Prime Intellect: INTELLECT 3"},
		{input: "Qwen: Qwen Plus 0728", expected: "Qwen: Qwen Plus 0728"},
		{input: "Qwen: Qwen Plus 0728 (thinking)", expected: "Qwen: Qwen Plus 0728 (thinking)"},
		{input: "Qwen: Qwen VL Max", expected: "Qwen: Qwen VL Max"},
		{input: "Qwen: Qwen VL Plus", expected: "Qwen: Qwen VL Plus"},
		{input: "Qwen: Qwen-Max", expected: "Qwen: Qwen Max"},
		{input: "Qwen: Qwen-Plus", expected: "Qwen: Qwen Plus"},
		{input: "Qwen: Qwen-Turbo", expected: "Qwen: Qwen Turbo"},
		{input: "Qwen: Qwen2.5 7B Instruct", expected: "Qwen: Qwen2.5 7B Instruct"},
		{input: "Qwen: Qwen2.5 Coder 7B Instruct", expected: "Qwen: Qwen2.5 Coder 7B Instruct"},
		{input: "Qwen: Qwen2.5 VL 32B Instruct", expected: "Qwen: Qwen2.5 VL 32B Instruct"},
		{input: "Qwen: Qwen2.5 VL 72B Instruct", expected: "Qwen: Qwen2.5 VL 72B Instruct"},
		{input: "Qwen: Qwen2.5-VL 7B Instruct", expected: "Qwen: Qwen2.5 VL 7B Instruct"},
		{input: "Qwen: Qwen2.5-VL 7B Instruct (free)", expected: "Qwen: Qwen2.5 VL 7B Instruct (free)"},
		{input: "Qwen: Qwen3 14B", expected: "Qwen: Qwen3 14B"},
		{input: "Qwen: Qwen3 235B A22B", expected: "Qwen: Qwen3 235B A22B"},
		{input: "Qwen: Qwen3 235B A22B Instruct 2507", expected: "Qwen: Qwen3 235B A22B Instruct 2507"},
		{input: "Qwen: Qwen3 235B A22B Thinking 2507", expected: "Qwen: Qwen3 235B A22B Thinking 2507"},
		{input: "Qwen: Qwen3 30B A3B", expected: "Qwen: Qwen3 30B A3B"},
		{input: "Qwen: Qwen3 30B A3B Instruct 2507", expected: "Qwen: Qwen3 30B A3B Instruct 2507"},
		{input: "Qwen: Qwen3 30B A3B Thinking 2507", expected: "Qwen: Qwen3 30B A3B Thinking 2507"},
		{input: "Qwen: Qwen3 32B", expected: "Qwen: Qwen3 32B"},
		{input: "Qwen: Qwen3 4B (free)", expected: "Qwen: Qwen3 4B (free)"},
		{input: "Qwen: Qwen3 8B", expected: "Qwen: Qwen3 8B"},
		{input: "Qwen: Qwen3 Coder 30B A3B Instruct", expected: "Qwen: Qwen3 Coder 30B A3B Instruct"},
		{input: "Qwen: Qwen3 Coder 480B A35B", expected: "Qwen: Qwen3 Coder 480B A35B"},
		{input: "Qwen: Qwen3 Coder 480B A35B (exacto)", expected: "Qwen: Qwen3 Coder 480B A35B (exacto)"},
		{input: "Qwen: Qwen3 Coder 480B A35B (free)", expected: "Qwen: Qwen3 Coder 480B A35B (free)"},
		{input: "Qwen: Qwen3 Coder Flash", expected: "Qwen: Qwen3 Coder Flash"},
		{input: "Qwen: Qwen3 Coder Plus", expected: "Qwen: Qwen3 Coder Plus"},
		{input: "Qwen: Qwen3 Max", expected: "Qwen: Qwen3 Max"},
		{input: "Qwen: Qwen3 Next 80B A3B Instruct", expected: "Qwen: Qwen3 Next 80B A3B Instruct"},
		{input: "Qwen: Qwen3 Next 80B A3B Thinking", expected: "Qwen: Qwen3 Next 80B A3B Thinking"},
		{input: "Qwen: Qwen3 VL 235B A22B Instruct", expected: "Qwen: Qwen3 VL 235B A22B Instruct"},
		{input: "Qwen: Qwen3 VL 235B A22B Thinking", expected: "Qwen: Qwen3 VL 235B A22B Thinking"},
		{input: "Qwen: Qwen3 VL 30B A3B Instruct", expected: "Qwen: Qwen3 VL 30B A3B Instruct"},
		{input: "Qwen: Qwen3 VL 30B A3B Thinking", expected: "Qwen: Qwen3 VL 30B A3B Thinking"},
		{input: "Qwen: Qwen3 VL 32B Instruct", expected: "Qwen: Qwen3 VL 32B Instruct"},
		{input: "Qwen: Qwen3 VL 8B Instruct", expected: "Qwen: Qwen3 VL 8B Instruct"},
		{input: "Qwen: Qwen3 VL 8B Thinking", expected: "Qwen: Qwen3 VL 8B Thinking"},
		{input: "Qwen: QwQ 32B", expected: "Qwen: QwQ 32B"},
		{input: "Qwen2.5 72B Instruct", expected: "Qwen2.5 72B Instruct"},
		{input: "Qwen2.5 Coder 32B Instruct", expected: "Qwen2.5 Coder 32B Instruct"},
		{input: "Relace: Relace Apply 3", expected: "Relace: Relace Apply 3"},
		{input: "Relace: Relace Search", expected: "Relace: Relace Search"},
		{input: "ReMM SLERP 13B", expected: "ReMM SLERP 13B"},
		{input: "Sao10K: Llama 3 8B Lunaris", expected: "Sao10k: Llama 3 8B Lunaris"},
		{input: "Sao10k: Llama 3 Euryale 70B v2.1", expected: "Sao10k: Llama 3 Euryale 70B v2.1"},
		{input: "Sao10K: Llama 3.1 70B Hanami x1", expected: "Sao10k: Llama 3.1 70B Hanami X1"},
		{input: "Sao10K: Llama 3.1 Euryale 70B v2.2", expected: "Sao10k: Llama 3.1 Euryale 70B v2.2"},
		{input: "Sao10K: Llama 3.3 Euryale 70B", expected: "Sao10k: Llama 3.3 Euryale 70B"},
		{input: "SorcererLM 8x22B", expected: "SorcererLM 8x22B"},
		{input: "StepFun: Step3", expected: "StepFun: Step3"},
		{input: "Switchpoint Router", expected: "Switchpoint Router"},
		{input: "Tencent: Hunyuan A13B Instruct", expected: "Tencent: Hunyuan A13B Instruct"},
		{input: "TheDrummer: Cydonia 24B V4.1", expected: "TheDrummer: Cydonia 24B v4.1"},
		{input: "TheDrummer: Rocinante 12B", expected: "TheDrummer: Rocinante 12B"},
		{input: "TheDrummer: Skyfall 36B V2", expected: "TheDrummer: Skyfall 36B v2"},
		{input: "TheDrummer: UnslopNemo 12B", expected: "TheDrummer: UnslopNemo 12B"},
		{input: "THUDM: GLM 4.1V 9B Thinking", expected: "THUDM: GLM 4.1V 9B Thinking"},
		{input: "TNG: DeepSeek R1T Chimera", expected: "TNG: DeepSeek R1T Chimera"},
		{input: "TNG: DeepSeek R1T Chimera (free)", expected: "TNG: DeepSeek R1T Chimera (free)"},
		{input: "TNG: DeepSeek R1T2 Chimera", expected: "TNG: DeepSeek R1T2 Chimera"},
		{input: "TNG: DeepSeek R1T2 Chimera (free)", expected: "TNG: DeepSeek R1T2 Chimera (free)"},
		{input: "TNG: R1T Chimera", expected: "TNG: R1T Chimera"},
		{input: "TNG: R1T Chimera (free)", expected: "TNG: R1T Chimera (free)"},
		{input: "Tongyi DeepResearch 30B A3B", expected: "Tongyi DeepResearch 30B A3B"},
		{input: "Tongyi DeepResearch 30B A3B (free)", expected: "Tongyi DeepResearch 30B A3B (free)"},
		{input: "Venice: Uncensored (free)", expected: "Venice: Uncensored (free)"},
		{input: "WizardLM-2 8x22B", expected: "WizardLM 2 8x22B"},
		{input: "xAI: Grok 3", expected: "xAI: Grok 3"},
		{input: "xAI: Grok 3 Beta", expected: "xAI: Grok 3 Beta"},
		{input: "xAI: Grok 3 Mini", expected: "xAI: Grok 3 Mini"},
		{input: "xAI: Grok 3 Mini Beta", expected: "xAI: Grok 3 Mini Beta"},
		{input: "xAI: Grok 4", expected: "xAI: Grok 4"},
		{input: "xAI: Grok 4 Fast", expected: "xAI: Grok 4 Fast"},
		{input: "xAI: Grok 4.1 Fast", expected: "xAI: Grok 4.1 Fast"},
		{input: "xAI: Grok Code Fast 1", expected: "xAI: Grok Code Fast 1"},
		{input: "Xiaomi: MiMo-V2-Flash (free)", expected: "Xiaomi: MiMo v2 Flash (free)"},
		{input: "Z.AI: GLM 4 32B", expected: "Z.AI: GLM 4 32B"},
		{input: "Z.AI: GLM 4.5", expected: "Z.AI: GLM 4.5"},
		{input: "Z.AI: GLM 4.5 Air", expected: "Z.AI: GLM 4.5 Air"},
		{input: "Z.AI: GLM 4.5 Air (free)", expected: "Z.AI: GLM 4.5 Air (free)"},
		{input: "Z.AI: GLM 4.5V", expected: "Z.AI: GLM 4.5V"},
		{input: "Z.AI: GLM 4.6", expected: "Z.AI: GLM 4.6"},
		{input: "Z.AI: GLM 4.6 (exacto)", expected: "Z.AI: GLM 4.6 (exacto)"},
		{input: "Z.AI: GLM 4.6V", expected: "Z.AI: GLM 4.6V"},
		{input: "o3 Pro", expected: "o3 Pro"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatModelTitle(tt.input))
		})
	}
}

// slugs
// curl https://openrouter.ai/api/v1/models  -H "Authorization: Bearer $OPENROUTER_API_KEY" | jq -r '.data[].id' | sort
// titles
// curl https://openrouter.ai/api/v1/models  -H "Authorization: Bearer $OPENROUTER_API_KEY" | jq -r '.data[].name' | sort
func TestPrintRealSlugsAndNamesAsTestCases(t *testing.T) {
	slugs := `ai21/jamba-large-1.7
ai21/jamba-mini-1.7
aion-labs/aion-1.0
aion-labs/aion-1.0-mini
aion-labs/aion-rp-llama-3.1-8b
alfredpros/codellama-7b-instruct-solidity
alibaba/tongyi-deepresearch-30b-a3b
alibaba/tongyi-deepresearch-30b-a3b:free
allenai/olmo-2-0325-32b-instruct
allenai/olmo-3-32b-think:free
allenai/olmo-3-7b-instruct
allenai/olmo-3-7b-think
allenai/olmo-3.1-32b-think:free
alpindale/goliath-120b
amazon/nova-2-lite-v1
amazon/nova-lite-v1
amazon/nova-micro-v1
amazon/nova-premier-v1
amazon/nova-pro-v1
anthracite-org/magnum-v4-72b
anthropic/claude-3-haiku
anthropic/claude-3-opus
anthropic/claude-3.5-haiku
anthropic/claude-3.5-haiku-20241022
anthropic/claude-3.5-sonnet
anthropic/claude-3.7-sonnet
anthropic/claude-3.7-sonnet:thinking
anthropic/claude-haiku-4.5
anthropic/claude-opus-4
anthropic/claude-opus-4.1
anthropic/claude-opus-4.5
anthropic/claude-sonnet-4
anthropic/claude-sonnet-4.5
arcee-ai/coder-large
arcee-ai/maestro-reasoning
arcee-ai/spotlight
arcee-ai/trinity-mini
arcee-ai/trinity-mini:free
arcee-ai/virtuoso-large
arliai/qwq-32b-arliai-rpr-v1
baidu/ernie-4.5-21b-a3b
baidu/ernie-4.5-21b-a3b-thinking
baidu/ernie-4.5-300b-a47b
baidu/ernie-4.5-vl-28b-a3b
baidu/ernie-4.5-vl-424b-a47b
bytedance/ui-tars-1.5-7b
cognitivecomputations/dolphin-mistral-24b-venice-edition:free
cohere/command-a
cohere/command-r-08-2024
cohere/command-r-plus-08-2024
cohere/command-r7b-12-2024
deepcogito/cogito-v2-preview-llama-109b-moe
deepcogito/cogito-v2-preview-llama-405b
deepcogito/cogito-v2-preview-llama-70b
deepcogito/cogito-v2.1-671b
deepseek/deepseek-chat
deepseek/deepseek-chat-v3-0324
deepseek/deepseek-chat-v3.1
deepseek/deepseek-prover-v2
deepseek/deepseek-r1
deepseek/deepseek-r1-0528
deepseek/deepseek-r1-0528-qwen3-8b
deepseek/deepseek-r1-0528:free
deepseek/deepseek-r1-distill-llama-70b
deepseek/deepseek-r1-distill-qwen-14b
deepseek/deepseek-r1-distill-qwen-32b
deepseek/deepseek-v3.1-terminus
deepseek/deepseek-v3.1-terminus:exacto
deepseek/deepseek-v3.2
deepseek/deepseek-v3.2-exp
deepseek/deepseek-v3.2-speciale
eleutherai/llemma_7b
essentialai/rnj-1-instruct
google/gemini-2.0-flash-001
google/gemini-2.0-flash-exp:free
google/gemini-2.0-flash-lite-001
google/gemini-2.5-flash
google/gemini-2.5-flash-image
google/gemini-2.5-flash-image-preview
google/gemini-2.5-flash-lite
google/gemini-2.5-flash-lite-preview-09-2025
google/gemini-2.5-flash-preview-09-2025
google/gemini-2.5-pro
google/gemini-2.5-pro-preview
google/gemini-2.5-pro-preview-05-06
google/gemini-3-flash-preview
google/gemini-3-pro-image-preview
google/gemini-3-pro-preview
google/gemma-2-27b-it
google/gemma-2-9b-it
google/gemma-3-12b-it
google/gemma-3-12b-it:free
google/gemma-3-27b-it
google/gemma-3-27b-it:free
google/gemma-3-4b-it
google/gemma-3-4b-it:free
google/gemma-3n-e2b-it:free
google/gemma-3n-e4b-it
google/gemma-3n-e4b-it:free
gryphe/mythomax-l2-13b
ibm-granite/granite-4.0-h-micro
inception/mercury
inception/mercury-coder
inflection/inflection-3-pi
inflection/inflection-3-productivity
kwaipilot/kat-coder-pro:free
liquid/lfm-2.2-6b
liquid/lfm2-8b-a1b
mancer/weaver
meituan/longcat-flash-chat
meta-llama/llama-3-70b-instruct
meta-llama/llama-3-8b-instruct
meta-llama/llama-3.1-405b
meta-llama/llama-3.1-405b-instruct
meta-llama/llama-3.1-405b-instruct:free
meta-llama/llama-3.1-70b-instruct
meta-llama/llama-3.1-8b-instruct
meta-llama/llama-3.2-11b-vision-instruct
meta-llama/llama-3.2-1b-instruct
meta-llama/llama-3.2-3b-instruct
meta-llama/llama-3.2-3b-instruct:free
meta-llama/llama-3.2-90b-vision-instruct
meta-llama/llama-3.3-70b-instruct
meta-llama/llama-3.3-70b-instruct:free
meta-llama/llama-4-maverick
meta-llama/llama-4-scout
meta-llama/llama-guard-2-8b
meta-llama/llama-guard-3-8b
meta-llama/llama-guard-4-12b
microsoft/phi-3-medium-128k-instruct
microsoft/phi-3-mini-128k-instruct
microsoft/phi-3.5-mini-128k-instruct
microsoft/phi-4
microsoft/phi-4-multimodal-instruct
microsoft/phi-4-reasoning-plus
microsoft/wizardlm-2-8x22b
minimax/minimax-01
minimax/minimax-m1
minimax/minimax-m2
mistralai/codestral-2508
mistralai/devstral-2512
mistralai/devstral-2512:free
mistralai/devstral-medium
mistralai/devstral-small
mistralai/devstral-small-2505
mistralai/ministral-14b-2512
mistralai/ministral-3b
mistralai/ministral-3b-2512
mistralai/ministral-8b
mistralai/ministral-8b-2512
mistralai/mistral-7b-instruct
mistralai/mistral-7b-instruct-v0.1
mistralai/mistral-7b-instruct-v0.2
mistralai/mistral-7b-instruct-v0.3
mistralai/mistral-7b-instruct:free
mistralai/mistral-large
mistralai/mistral-large-2407
mistralai/mistral-large-2411
mistralai/mistral-large-2512
mistralai/mistral-medium-3
mistralai/mistral-medium-3.1
mistralai/mistral-nemo
mistralai/mistral-saba
mistralai/mistral-small-24b-instruct-2501
mistralai/mistral-small-3.1-24b-instruct
mistralai/mistral-small-3.1-24b-instruct:free
mistralai/mistral-small-3.2-24b-instruct
mistralai/mistral-small-creative
mistralai/mistral-tiny
mistralai/mixtral-8x22b-instruct
mistralai/mixtral-8x7b-instruct
mistralai/pixtral-12b
mistralai/pixtral-large-2411
mistralai/voxtral-small-24b-2507
moonshotai/kimi-dev-72b
moonshotai/kimi-k2
moonshotai/kimi-k2-0905
moonshotai/kimi-k2-0905:exacto
moonshotai/kimi-k2-thinking
moonshotai/kimi-k2:free
morph/morph-v3-fast
morph/morph-v3-large
neversleep/llama-3.1-lumimaid-8b
neversleep/noromaid-20b
nex-agi/deepseek-v3.1-nex-n1:free
nousresearch/deephermes-3-mistral-24b-preview
nousresearch/hermes-2-pro-llama-3-8b
nousresearch/hermes-3-llama-3.1-405b
nousresearch/hermes-3-llama-3.1-405b:free
nousresearch/hermes-3-llama-3.1-70b
nousresearch/hermes-4-405b
nousresearch/hermes-4-70b
nvidia/llama-3.1-nemotron-70b-instruct
nvidia/llama-3.1-nemotron-ultra-253b-v1
nvidia/llama-3.3-nemotron-super-49b-v1.5
nvidia/nemotron-3-nano-30b-a3b
nvidia/nemotron-3-nano-30b-a3b:free
nvidia/nemotron-nano-12b-v2-vl
nvidia/nemotron-nano-12b-v2-vl:free
nvidia/nemotron-nano-9b-v2
nvidia/nemotron-nano-9b-v2:free
openai/chatgpt-4o-latest
openai/codex-mini
openai/gpt-3.5-turbo
openai/gpt-3.5-turbo-0613
openai/gpt-3.5-turbo-16k
openai/gpt-3.5-turbo-instruct
openai/gpt-4
openai/gpt-4-0314
openai/gpt-4-1106-preview
openai/gpt-4-turbo
openai/gpt-4-turbo-preview
openai/gpt-4.1
openai/gpt-4.1-mini
openai/gpt-4.1-nano
openai/gpt-4o
openai/gpt-4o-2024-05-13
openai/gpt-4o-2024-08-06
openai/gpt-4o-2024-11-20
openai/gpt-4o-audio-preview
openai/gpt-4o-mini
openai/gpt-4o-mini-2024-07-18
openai/gpt-4o-mini-search-preview
openai/gpt-4o-search-preview
openai/gpt-4o:extended
openai/gpt-5
openai/gpt-5-chat
openai/gpt-5-codex
openai/gpt-5-image
openai/gpt-5-image-mini
openai/gpt-5-mini
openai/gpt-5-nano
openai/gpt-5-pro
openai/gpt-5.1
openai/gpt-5.1-chat
openai/gpt-5.1-codex
openai/gpt-5.1-codex-max
openai/gpt-5.1-codex-mini
openai/gpt-5.2
openai/gpt-5.2-chat
openai/gpt-5.2-pro
openai/gpt-oss-120b
openai/gpt-oss-120b:exacto
openai/gpt-oss-120b:free
openai/gpt-oss-20b
openai/gpt-oss-20b:free
openai/gpt-oss-safeguard-20b
openai/o1
openai/o1-pro
openai/o3
openai/o3-deep-research
openai/o3-mini
openai/o3-mini-high
openai/o3-pro
openai/o4-mini
openai/o4-mini-deep-research
openai/o4-mini-high
opengvlab/internvl3-78b
openrouter/auto
openrouter/bodybuilder
perplexity/sonar
perplexity/sonar-deep-research
perplexity/sonar-pro
perplexity/sonar-pro-search
perplexity/sonar-reasoning
perplexity/sonar-reasoning-pro
prime-intellect/intellect-3
qwen/qwen-2.5-72b-instruct
qwen/qwen-2.5-7b-instruct
qwen/qwen-2.5-coder-32b-instruct
qwen/qwen-2.5-vl-7b-instruct
qwen/qwen-2.5-vl-7b-instruct:free
qwen/qwen-max
qwen/qwen-plus
qwen/qwen-plus-2025-07-28
qwen/qwen-plus-2025-07-28:thinking
qwen/qwen-turbo
qwen/qwen-vl-max
qwen/qwen-vl-plus
qwen/qwen2.5-coder-7b-instruct
qwen/qwen2.5-vl-32b-instruct
qwen/qwen2.5-vl-72b-instruct
qwen/qwen3-14b
qwen/qwen3-235b-a22b
qwen/qwen3-235b-a22b-2507
qwen/qwen3-235b-a22b-thinking-2507
qwen/qwen3-30b-a3b
qwen/qwen3-30b-a3b-instruct-2507
qwen/qwen3-30b-a3b-thinking-2507
qwen/qwen3-32b
qwen/qwen3-4b:free
qwen/qwen3-8b
qwen/qwen3-coder
qwen/qwen3-coder-30b-a3b-instruct
qwen/qwen3-coder-flash
qwen/qwen3-coder-plus
qwen/qwen3-coder:exacto
qwen/qwen3-coder:free
qwen/qwen3-max
qwen/qwen3-next-80b-a3b-instruct
qwen/qwen3-next-80b-a3b-thinking
qwen/qwen3-vl-235b-a22b-instruct
qwen/qwen3-vl-235b-a22b-thinking
qwen/qwen3-vl-30b-a3b-instruct
qwen/qwen3-vl-30b-a3b-thinking
qwen/qwen3-vl-32b-instruct
qwen/qwen3-vl-8b-instruct
qwen/qwen3-vl-8b-thinking
qwen/qwq-32b
raifle/sorcererlm-8x22b
relace/relace-apply-3
relace/relace-search
sao10k/l3-euryale-70b
sao10k/l3-lunaris-8b
sao10k/l3.1-70b-hanami-x1
sao10k/l3.1-euryale-70b
sao10k/l3.3-euryale-70b
stepfun-ai/step3
switchpoint/router
tencent/hunyuan-a13b-instruct
thedrummer/cydonia-24b-v4.1
thedrummer/rocinante-12b
thedrummer/skyfall-36b-v2
thedrummer/unslopnemo-12b
thudm/glm-4.1v-9b-thinking
tngtech/deepseek-r1t-chimera
tngtech/deepseek-r1t-chimera:free
tngtech/deepseek-r1t2-chimera
tngtech/deepseek-r1t2-chimera:free
tngtech/tng-r1t-chimera
tngtech/tng-r1t-chimera:free
undi95/remm-slerp-l2-13b
x-ai/grok-3
x-ai/grok-3-beta
x-ai/grok-3-mini
x-ai/grok-3-mini-beta
x-ai/grok-4
x-ai/grok-4-fast
x-ai/grok-4.1-fast
x-ai/grok-code-fast-1
xiaomi/mimo-v2-flash:free
z-ai/glm-4-32b
z-ai/glm-4.5
z-ai/glm-4.5-air
z-ai/glm-4.5-air:free
z-ai/glm-4.5v
z-ai/glm-4.6
z-ai/glm-4.6:exacto
z-ai/glm-4.6v`

	titles := `AI21: Jamba Large 1.7
AI21: Jamba Mini 1.7
AionLabs: Aion-1.0
AionLabs: Aion-1.0-Mini
AionLabs: Aion-RP 1.0 (8B)
AlfredPros: CodeLLaMa 7B Instruct Solidity
AllenAI: Olmo 2 32B Instruct
AllenAI: Olmo 3 32B Think (free)
AllenAI: Olmo 3 7B Instruct
AllenAI: Olmo 3 7B Think
AllenAI: Olmo 3.1 32B Think (free)
Amazon: Nova 2 Lite
Amazon: Nova Lite 1.0
Amazon: Nova Micro 1.0
Amazon: Nova Premier 1.0
Amazon: Nova Pro 1.0
Anthropic: Claude 3 Haiku
Anthropic: Claude 3 Opus
Anthropic: Claude 3.5 Haiku
Anthropic: Claude 3.5 Haiku (2024-10-22)
Anthropic: Claude 3.5 Sonnet
Anthropic: Claude 3.7 Sonnet
Anthropic: Claude 3.7 Sonnet (thinking)
Anthropic: Claude Haiku 4.5
Anthropic: Claude Opus 4
Anthropic: Claude Opus 4.1
Anthropic: Claude Opus 4.5
Anthropic: Claude Sonnet 4
Anthropic: Claude Sonnet 4.5
ArceeAI: Coder Large
ArceeAI: Maestro Reasoning
ArceeAI: Spotlight
ArceeAI: Trinity Mini
ArceeAI: Trinity Mini (free)
ArceeAI: Virtuoso Large
ArliAI: QwQ 32B RpR v1
Auto Router
Baidu: ERNIE 4.5 21B A3B
Baidu: ERNIE 4.5 21B A3B Thinking
Baidu: ERNIE 4.5 300B A47B
Baidu: ERNIE 4.5 VL 28B A3B
Baidu: ERNIE 4.5 VL 424B A47B
ByteDance: UI-TARS 7B
Body Builder (beta)
Cogito V2 Preview Llama 109B
Cohere: Command A
Cohere: Command R (08-2024)
Cohere: Command R+ (08-2024)
Cohere: Command R7B (12-2024)
Deep Cogito: Cogito V2 Preview Llama 405B
Deep Cogito: Cogito V2 Preview Llama 70B
Deep Cogito: Cogito v2.1 671B
DeepSeek: DeepSeek Prover V2
DeepSeek: DeepSeek R1 0528 Qwen3 8B
DeepSeek: DeepSeek V3
DeepSeek: DeepSeek V3 0324
DeepSeek: DeepSeek V3.1
DeepSeek: DeepSeek V3.1 Terminus
DeepSeek: DeepSeek V3.1 Terminus (exacto)
DeepSeek: DeepSeek V3.2
DeepSeek: DeepSeek V3.2 Exp
DeepSeek: DeepSeek V3.2 Speciale
DeepSeek: R1
DeepSeek: R1 0528
DeepSeek: R1 0528 (free)
DeepSeek: R1 Distill Llama 70B
DeepSeek: R1 Distill Qwen 14B
DeepSeek: R1 Distill Qwen 32B
EleutherAI: Llemma 7b
EssentialAI: Rnj 1 Instruct
Goliath 120B
Google: Gemini 2.0 Flash
Google: Gemini 2.0 Flash Experimental (free)
Google: Gemini 2.0 Flash Lite
Google: Gemini 2.5 Flash
Google: Gemini 2.5 Flash Image (Nano Banana)
Google: Gemini 2.5 Flash Image Preview (Nano Banana)
Google: Gemini 2.5 Flash Lite
Google: Gemini 2.5 Flash Lite Preview 09-2025
Google: Gemini 2.5 Flash Preview 09-2025
Google: Gemini 2.5 Pro
Google: Gemini 2.5 Pro Preview 05-06
Google: Gemini 2.5 Pro Preview 06-05
Google: Gemini 3 Flash Preview
Google: Gemini 3 Pro Preview
Google: Gemma 2 27B
Google: Gemma 2 9B
Google: Gemma 3 12B
Google: Gemma 3 12B (free)
Google: Gemma 3 27B
Google: Gemma 3 27B (free)
Google: Gemma 3 4B
Google: Gemma 3 4B (free)
Google: Gemma 3n 2B (free)
Google: Gemma 3n 4B
Google: Gemma 3n 4B (free)
Google: Nano Banana Pro (Gemini 3 Pro Image Preview)
IBM: Granite 4.0 Micro
Inception: Mercury
Inception: Mercury Coder
Inflection: Inflection 3 Pi
Inflection: Inflection 3 Productivity
Kwaipilot: KAT-Coder-Pro V1 (free)
LiquidAI/LFM2-2.6B
LiquidAI/LFM2-8B-A1B
Llama Guard 3 8B
Magnum v4 72B
Mancer: Weaver (alpha)
Meituan: LongCat Flash Chat
Meta: Llama 3 70B Instruct
Meta: Llama 3 8B Instruct
Meta: Llama 3.1 405B (base)
Meta: Llama 3.1 405B Instruct
Meta: Llama 3.1 405B Instruct (free)
Meta: Llama 3.1 70B Instruct
Meta: Llama 3.1 8B Instruct
Meta: Llama 3.2 11B Vision Instruct
Meta: Llama 3.2 1B Instruct
Meta: Llama 3.2 3B Instruct
Meta: Llama 3.2 3B Instruct (free)
Meta: Llama 3.2 90B Vision Instruct
Meta: Llama 3.3 70B Instruct
Meta: Llama 3.3 70B Instruct (free)
Meta: Llama 4 Maverick
Meta: Llama 4 Scout
Meta: Llama Guard 4 12B
Meta: LlamaGuard 2 8B
Microsoft: Phi 4
Microsoft: Phi 4 Multimodal Instruct
Microsoft: Phi 4 Reasoning Plus
Microsoft: Phi-3 Medium 128K Instruct
Microsoft: Phi-3 Mini 128K Instruct
Microsoft: Phi-3.5 Mini 128K Instruct
MiniMax: MiniMax M1
MiniMax: MiniMax M2
MiniMax: MiniMax-01
Mistral Large
Mistral Large 2407
Mistral Large 2411
Mistral Tiny
Mistral: Codestral 2508
Mistral: Devstral 2 2512
Mistral: Devstral 2 2512 (free)
Mistral: Devstral Medium
Mistral: Devstral Small 1.1
Mistral: Devstral Small 2505
Mistral: Ministral 3 14B 2512
Mistral: Ministral 3 3B 2512
Mistral: Ministral 3 8B 2512
Mistral: Ministral 3B
Mistral: Ministral 8B
Mistral: Mistral 7B Instruct
Mistral: Mistral 7B Instruct (free)
Mistral: Mistral 7B Instruct v0.1
Mistral: Mistral 7B Instruct v0.2
Mistral: Mistral 7B Instruct v0.3
Mistral: Mistral Large 3 2512
Mistral: Mistral Medium 3
Mistral: Mistral Medium 3.1
Mistral: Mistral Nemo
Mistral: Mistral Small 3
Mistral: Mistral Small 3.1 24B
Mistral: Mistral Small 3.1 24B (free)
Mistral: Mistral Small 3.2 24B
Mistral: Mistral Small Creative
Mistral: Mixtral 8x22B Instruct
Mistral: Mixtral 8x7B Instruct
Mistral: Pixtral 12B
Mistral: Pixtral Large 2411
Mistral: Saba
Mistral: Voxtral Small 24B 2507
MythoMax 13B
MoonshotAI: Kimi Dev 72B
MoonshotAI: Kimi K2 0711
MoonshotAI: Kimi K2 0711 (free)
MoonshotAI: Kimi K2 0905
MoonshotAI: Kimi K2 0905 (exacto)
MoonshotAI: Kimi K2 Thinking
Morph: Morph V3 Fast
Morph: Morph V3 Large
NeverSleep: Lumimaid v0.2 8B
Nex AGI: DeepSeek V3.1 Nex N1 (free)
Noromaid 20B
Nous: DeepHermes 3 Mistral 24B Preview
Nous: Hermes 3 405B Instruct
Nous: Hermes 3 405B Instruct (free)
Nous: Hermes 3 70B Instruct
Nous: Hermes 4 405B
Nous: Hermes 4 70B
NousResearch: Hermes 2 Pro - Llama-3 8B
NVIDIA: Llama 3.1 Nemotron 70B Instruct
NVIDIA: Llama 3.1 Nemotron Ultra 253B v1
NVIDIA: Llama 3.3 Nemotron Super 49B V1.5
NVIDIA: Nemotron 3 Nano 30B A3B
NVIDIA: Nemotron 3 Nano 30B A3B (free)
NVIDIA: Nemotron Nano 12B 2 VL
NVIDIA: Nemotron Nano 12B 2 VL (free)
NVIDIA: Nemotron Nano 9B V2
NVIDIA: Nemotron Nano 9B V2 (free)
OpenAI: ChatGPT-4o
OpenAI: Codex Mini
OpenAI: GPT-3.5 Turbo
OpenAI: GPT-3.5 Turbo (older v0613)
OpenAI: GPT-3.5 Turbo 16k
OpenAI: GPT-3.5 Turbo Instruct
OpenAI: GPT-4
OpenAI: GPT-4 (older v0314)
OpenAI: GPT-4 Turbo
OpenAI: GPT-4 Turbo (older v1106)
OpenAI: GPT-4 Turbo Preview
OpenAI: GPT-4.1
OpenAI: GPT-4.1 Mini
OpenAI: GPT-4.1 Nano
OpenAI: GPT-4o
OpenAI: GPT-4o (2024-05-13)
OpenAI: GPT-4o (2024-08-06)
OpenAI: GPT-4o (2024-11-20)
OpenAI: GPT-4o (extended)
OpenAI: GPT-4o Audio
OpenAI: GPT-4o Search Preview
OpenAI: GPT-4o-mini
OpenAI: GPT-4o-mini (2024-07-18)
OpenAI: GPT-4o-mini Search Preview
OpenAI: GPT-5
OpenAI: GPT-5 Chat
OpenAI: GPT-5 Codex
OpenAI: GPT-5 Image
OpenAI: GPT-5 Image Mini
OpenAI: GPT-5 Mini
OpenAI: GPT-5 Nano
OpenAI: GPT-5 Pro
OpenAI: GPT-5.1
OpenAI: GPT-5.1 Chat
OpenAI: GPT-5.1-Codex
OpenAI: GPT-5.1-Codex-Max
OpenAI: GPT-5.1-Codex-Mini
OpenAI: GPT-5.2
OpenAI: GPT-5.2 Chat
OpenAI: GPT-5.2 Pro
OpenAI: gpt-oss-120b
OpenAI: gpt-oss-120b (exacto)
OpenAI: gpt-oss-120b (free)
OpenAI: gpt-oss-20b
OpenAI: gpt-oss-20b (free)
OpenAI: gpt-oss-safeguard-20b
OpenAI: o1
OpenAI: o1-pro
OpenAI: o3
OpenAI: o3 Deep Research
OpenAI: o3 Mini
OpenAI: o3 Mini High
OpenAI: o3 Pro
OpenAI: o4 Mini
OpenAI: o4 Mini Deep Research
OpenAI: o4 Mini High
OpenGVLab: InternVL3 78B
Perplexity: Sonar
Perplexity: Sonar Deep Research
Perplexity: Sonar Pro
Perplexity: Sonar Pro Search
Perplexity: Sonar Reasoning
Perplexity: Sonar Reasoning Pro
Prime Intellect: INTELLECT-3
Qwen: Qwen Plus 0728
Qwen: Qwen Plus 0728 (thinking)
Qwen: Qwen VL Max
Qwen: Qwen VL Plus
Qwen: Qwen-Max
Qwen: Qwen-Plus
Qwen: Qwen-Turbo
Qwen: Qwen2.5 7B Instruct
Qwen: Qwen2.5 Coder 7B Instruct
Qwen: Qwen2.5 VL 32B Instruct
Qwen: Qwen2.5 VL 72B Instruct
Qwen: Qwen2.5-VL 7B Instruct
Qwen: Qwen2.5-VL 7B Instruct (free)
Qwen: Qwen3 14B
Qwen: Qwen3 235B A22B
Qwen: Qwen3 235B A22B Instruct 2507
Qwen: Qwen3 235B A22B Thinking 2507
Qwen: Qwen3 30B A3B
Qwen: Qwen3 30B A3B Instruct 2507
Qwen: Qwen3 30B A3B Thinking 2507
Qwen: Qwen3 32B
Qwen: Qwen3 4B (free)
Qwen: Qwen3 8B
Qwen: Qwen3 Coder 30B A3B Instruct
Qwen: Qwen3 Coder 480B A35B
Qwen: Qwen3 Coder 480B A35B (exacto)
Qwen: Qwen3 Coder 480B A35B (free)
Qwen: Qwen3 Coder Flash
Qwen: Qwen3 Coder Plus
Qwen: Qwen3 Max
Qwen: Qwen3 Next 80B A3B Instruct
Qwen: Qwen3 Next 80B A3B Thinking
Qwen: Qwen3 VL 235B A22B Instruct
Qwen: Qwen3 VL 235B A22B Thinking
Qwen: Qwen3 VL 30B A3B Instruct
Qwen: Qwen3 VL 30B A3B Thinking
Qwen: Qwen3 VL 32B Instruct
Qwen: Qwen3 VL 8B Instruct
Qwen: Qwen3 VL 8B Thinking
Qwen: QwQ 32B
Qwen2.5 72B Instruct
Qwen2.5 Coder 32B Instruct
Relace: Relace Apply 3
Relace: Relace Search
ReMM SLERP 13B
Sao10K: Llama 3 8B Lunaris
Sao10k: Llama 3 Euryale 70B v2.1
Sao10K: Llama 3.1 70B Hanami x1
Sao10K: Llama 3.1 Euryale 70B v2.2
Sao10K: Llama 3.3 Euryale 70B
SorcererLM 8x22B
StepFun: Step3
Switchpoint Router
Tencent: Hunyuan A13B Instruct
TheDrummer: Cydonia 24B V4.1
TheDrummer: Rocinante 12B
TheDrummer: Skyfall 36B V2
TheDrummer: UnslopNemo 12B
THUDM: GLM 4.1V 9B Thinking
TNG: DeepSeek R1T Chimera
TNG: DeepSeek R1T Chimera (free)
TNG: DeepSeek R1T2 Chimera
TNG: DeepSeek R1T2 Chimera (free)
TNG: R1T Chimera
TNG: R1T Chimera (free)
Tongyi DeepResearch 30B A3B
Tongyi DeepResearch 30B A3B (free)
Venice: Uncensored (free)
WizardLM-2 8x22B
xAI: Grok 3
xAI: Grok 3 Beta
xAI: Grok 3 Mini
xAI: Grok 3 Mini Beta
xAI: Grok 4
xAI: Grok 4 Fast
xAI: Grok 4.1 Fast
xAI: Grok Code Fast 1
Xiaomi: MiMo-V2-Flash (free)
Z.AI: GLM 4 32B
Z.AI: GLM 4.5
Z.AI: GLM 4.5 Air
Z.AI: GLM 4.5 Air (free)
Z.AI: GLM 4.5V
Z.AI: GLM 4.6
Z.AI: GLM 4.6 (exacto)
Z.AI: GLM 4.6V`

	_ = slugs
	_ = titles

	// fields := strings.Split(slugs, "\n")
	// for _, field := range fields {
	// 	fmt.Printf(`{input: "%s", expected: "%s"},`+"\n", field, FormatModelTitle(field))
	// }
}
