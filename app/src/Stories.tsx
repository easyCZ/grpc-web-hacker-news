import * as React from 'react';
import { connect, Dispatch } from 'react-redux';
import { RootState } from './store';
import { Container, Grid, Header } from 'semantic-ui-react';
import StoryList from './StoryList';
import StoryView from './StoryView';
import { RootAction } from './actions';
import { listStories, selectStory } from './actions/stories';
import { Story } from './proto/hackernews_pb';

type StoriesProps = {
  stories: Story.AsObject[],
  loading: boolean,
  error: Error | null,
  dispatch: Dispatch<RootAction>,
  selected: Story.AsObject | null,
};

class Stories extends React.Component<StoriesProps, {}> {

  constructor(props: StoriesProps) {
    super(props);
    this.state = {
      selected: null,
    };
  }

  componentDidMount() {
    this.props.dispatch(listStories());
  }

  render() {
    return (
      <Container style={{padding: '1em'}} fluid={true}>
        <Header as="h1" dividing={true}>Hacker News with gRPC-Web</Header>

        <Grid columns={2} stackable={true} divided={'vertically'}>
          <Grid.Column width={4}>
            <StoryList
              selected={this.props.selected}
              stories={this.props.stories}
              onStorySelect={(id: number) => this.props.dispatch(selectStory(id))}
            />
          </Grid.Column>

          <Grid.Column width={12} stretched={true}>
            { this.props.selected
              ? <StoryView story={this.props.selected} />
              : null
            }
          </Grid.Column>
        </Grid>

      </Container>
    );
  }

}

function mapStateToProps(state: RootState) {
  return {
    stories: Object.keys(state.stories.stories).map(key => state.stories.stories[key]),
    loading: state.stories.loading,
    error: state.stories.error,
    selected: state.stories.selected,
  };
}

export default connect(mapStateToProps)(Stories);
