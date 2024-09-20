<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>CLI Application Documentation</title>
        <style>
        body {
        font-family: Arial, sans-serif;
        line-height: 1.6;
        max-width: 800px;
        margin: 0 auto;
        padding: 20px;
        }
        h1, h2, h3 {
        color: #333;
        }
        code {
        background-color: #f4f4f4;
        padding: 5px;
        border-radius: 4px;
        }
        pre {
        background-color: #f4f4f4;
        padding: 10px;
        border-radius: 4px;
        overflow-x: auto;
        }
        .command {
        color: #008080;
        font-weight: bold;
        }
        .description {
        color: #555;
        }
        .example {
        background-color: #eef;
        padding: 10px;
        margin-top: 10px;
        border-left: 5px solid #008080;
        }
        </style>
    </head>
    <body>
        <h1>MyApp CLI Documentation</h1>
        <p><strong>MyApp</strong> is a command-line tool built with Go that helps manage tasks efficiently.</p>

        <h2>Installation</h2>
        <p>To install <code>MyApp</code>, run the following command:</p>
        <pre><code>go install github.com/yourusername/myapp</code></pre>

        <h2>Usage</h2>
        <p>Here are some basic commands you can use with <code>MyApp</code>:</p>

        <h3><span class="command">myapp add</span></h3>
        <p class="description">Adds a new task to the list.</p>
        <pre><code>myapp add "Complete the report"</code></pre>

        <h3><span class="command">myapp list</span></h3>
        <p class="description">Lists all tasks in the current task list.</p>
        <pre><code>myapp list</code></pre>

        <h3><span class="command">myapp remove [id]</span></h3>
        <p class="description">Removes a task by ID.</p>
        <pre><code>myapp remove 1</code></pre>

        <div class="example">
            <h3>Example</h3>
            <pre><code># Add a new task
                myapp add "Finish Go tutorial"

                # List all tasks
                myapp list

                # Remove a task by its ID
                myapp remove 1</code></pre>
        </div>

        <h2>Configuration</h2>
        <p><code>MyApp</code> uses a configuration file located in <code>~/.myapp/config.yaml</code> for settings. Make sure to set up the config file before running the app.</p>

        <h2>Contributing</h2>
        <p>If you'd like to contribute to the project, feel free to submit a pull request on GitHub!</p>

        <h2>License</h2>
        <p>This project is licensed under the MIT License.</p>
    </body>
</html>


