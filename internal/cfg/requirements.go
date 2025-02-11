package cfg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const test = `# Main dependencies
requests==2.25.0           # HTTP library
flask>=1.1.2,<2.0          # Web framework
numpy==1.19.2              # Numerical computations
pandas>=1.1.3,<1.3.0       # Data manipulation
scipy>=1.5.0               # Scientific computing
matplotlib==3.3.2          # Plotting library

# Development dependencies
pytest==6.1.2              # Testing framework
black==20.8b1              # Code formatting
flake8==3.8.4              # Linting tool
mypy==0.812                # Type checking

# Database dependencies
SQLAlchemy==1.3.23         # ORM for database interaction
psycopg2>=2.9.1            # PostgreSQL adapter
mysqlclient==2.1.0         # MySQL adapter

# For machine learning
tensorflow>=2.4,<2.6       # Deep learning framework
scikit-learn==0.24.0       # Machine learning library
xgboost==1.3.3             # Gradient boosting framework

# Extras for web
flask[async]==2.0.0        # Flask with async support
django[postgresql]==3.2.5  # Django with PostgreSQL support

# For handling data files
openpyxl>=3.0.0            # Excel file handling
python-dotenv==0.15.0      # Load environment variables from .env files

# Version pinning examples
pyyaml==5.4.1              # YAML handling
requests[security]==2.25.0 # Requests with security extras
`

type pydep struct {
	Name    string
	Version string
}

func ParseRequirements() {
	panic("WIP")
	file, err := os.Open("requirements.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	deps := []pydep{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		var name string
		var version string
		left := true
		for _, ch := range line {
			if ch == '#' {
				break
			} else {
				if ch == '=' || ch == '>' || ch == '!' {
					left = false
				}
				if ch == ';' {
					left = false
					continue
				}
				if left {
					name += string(ch)
				} else {
					version += string(ch)
				}
			}
		}
		name = strings.TrimSpace(name)
		version = strings.TrimSpace(version)
		if len(name) > 0 || len(version) > 0 {
			version = strings.TrimPrefix(version, "==")
			deps = append(deps, pydep{Name: name, Version: version})
		}
	}
	fmt.Println(deps)
}
