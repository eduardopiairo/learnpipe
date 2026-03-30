# AI Agents

## What is an Agent?

- An AI model capable of reasoning, planning, and interacting with its environment.
- We call it Agent because it has agency, aka it has the ability to interact with the environment.

>> An Agent is a system that leverages an AI model to interact with its environment in order to achieve a user-defined objective. It combines reasoning, planning, and the execution of actions (often via external tools) to fulfill tasks.

The Agent has two main parts:
1. The brain (AI Model) 
    - The AI model does the reasoning and planning.
    - It decides which actions to take based on the situation and the objective.
2. The body (Capabilities and Tools):
    - The tools and interfaces that allow the Agent to interact with the environment.


What type of tasks can an Agent do?

- The design of the Tools is very important and has a great impact on the quality of your Agent. Some tasks will require very specific Tools to be crafted, while others may be solved with general purpose tools like “web_search”.


To summarize, an Agent is a system that uses an AI Model (typically an LLM) as its core reasoning engine, to:
- Understand natural language: Interpret and respond to human instructions in a meaningful way.
- Reason and plan: Analyze information, make decisions, and devise strategies to solve problems.
- Interact with its environment: Gather information, take actions, and observe the results of those actions.

## What are LLMs?

What is a Large Language Model?
An LLM is a type of AI model that excels at understanding and generating human language. They are trained on vast amounts of text data, allowing them to learn patterns, structure, and even nuance in language. These models typically consist of many millions of parameters.

Attention is all you need.
A key aspect of the Transformer architecture is Attention. When predicting the next word, not every word in a sentence is equally important; words like “France” and “capital” in the sentence “The capital of France is …” carry the most meaning.

### Messages: the underlying system of LLMs

#### System Messages

- System messages (also called System Prompts) define how the model should behave. They serve as persistent instructions, guiding every subsequent interaction.
- When using Agents, the System Message also gives information about the available tools, provides instructions to the model on how to format the actions to take, and includes guidelines on how the thought process should be segmented.

#### Conversations: User and Assistant Messages

- A conversation consists of alternating messages between a Human (user) and an LLM (assistant).
- Chat templates help maintain context by preserving conversation history, storing previous exchanges between the user and the assistant. This leads to more coherent multi-turn conversations.

### Chat Templates

As mentioned, chat templates are essential for structuring conversations between language models and users. They guide how message exchanges are formatted into a single prompt.

#### Base Models vs. Instruct Models

Another point we need to understand is the difference between a Base Model vs. an Instruct Model:
- A Base Model is trained on raw text data to predict the next token.
- An Instruct Model is fine-tuned specifically to follow instructions and engage in conversations.

#### Understanding Chat Templates

- Because each instruct model uses different conversation formats and special tokens, chat templates are implemented to ensure that we correctly format the prompt the way each model expects.
- This structure helps maintain consistency across interactions and ensures the model responds appropriately to different types of inputs.

#### Messages to prompt

The easiest way to ensure your LLM receives a conversation correctly formatted is to use the *chat_template* from the model’s tokenizer.


## What are Tools?

A Tool is a function given to the LLM. This function should fulfill a clear objective.

Here are some commonly used tools in AI agents:
- Web Search
    - Allows the agent to fetch up-to-date information from the internet.
- Image Generation
    - Creates images based on text descriptions.
- Retrieval
    - Retrieves information from an external source.
- API Interface
    - Interacts with an external API (GitHub, YouTube, Spotify, etc.).

A good tool should be something that complements the power of an LLM.

LLMs predict the completion of a prompt based on their training data, which means that their internal knowledge only includes events prior to their training. Therefore, if your agent needs up-to-date data you must provide it through some tool.

### How do Tools work?


#### Model Context Protocol (MCP): a unified tool interface

Model Context Protocol (MCP) is an open protocol that standardizes how applications provide tools to LLMs.

MCP provides:
- A growing list of pre-built integrations that your LLM can directly plug into
- The flexibility to switch between LLM providers and vendors
- Best practices for securing your data within your infrastructure